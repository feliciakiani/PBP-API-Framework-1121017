package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(c *gin.Context) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM movie")
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Movie query"})
		return
	}

	var movie Movie
	for rows.Next() {
		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Duration, &movie.Language); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "movies not found"})
		} else {
			c.IndentedJSON(http.StatusOK, movie)
		}
	}
}

func GetMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	idMovie := c.Query("ID")
	fmt.Print("Function GetMovieNormal, idMovie = ", idMovie)

	row, err := db.Query("SELECT * FROM movie WHERE id=?", idMovie)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Something has gone wrong with the Movie query"})
		return
	}

	var movie Movie
	for row.Next() {
		if err := row.Scan(&movie.ID, &movie.Title, &movie.Duration, &movie.Language); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "movies not found"})
		} else {
			c.IndentedJSON(http.StatusOK, movie)
		}
	}

}

func InsertMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	var movie Movie

	if err := c.Bind(&movie); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err := db.Query("INSERT INTO movie (title, duration, language) VALUES (?,?,?)",
		movie.Title,
		movie.Duration,
		movie.Language,
	)

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "insert failed"})
	} else {

		c.IndentedJSON(http.StatusCreated, movie)
	}

}

func UpdateMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	var movie Movie

	if err := c.Bind(&movie); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(movie.ID, movie.Title, movie.Duration, movie.Language)

	result, err := db.Exec("UPDATE movie SET title=?, duration=?, language=? WHERE id=?",
		movie.Title,
		movie.Duration,
		movie.Language,
		movie.ID,
	)

	fmt.Println("result : ", result)
	fmt.Println("err : ", err)

	num, _ := result.RowsAffected()

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "update failed"})
	} else {
		if num == 0 {
			log.Println(err)
			c.JSON(400, gin.H{"error": "update failed - no rows affected"})
		} else {
			c.IndentedJSON(http.StatusCreated, movie)
		}
	}
}

func DeleteMovie(c *gin.Context) {
	db := connect()
	defer db.Close()

	var movie Movie
	idMovie := c.Query("ID")
	fmt.Println("idMovie = ", idMovie)

	if err := c.Bind(&movie); err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	result, err := db.Exec("DELETE FROM movie WHERE id=?", idMovie)

	num, _ := result.RowsAffected()

	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "delete failed"})
	} else {
		if num == 0 {
			log.Println(err)
			c.JSON(400, gin.H{"error": "delete failed - no rows affected "})
		} else {
			c.IndentedJSON(200, gin.H{"success": "delete success"})
		}

	}
}
