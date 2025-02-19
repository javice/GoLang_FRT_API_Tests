// tests/api/posts_test.go

package tests

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	"testing"
	"time"
	"log"
	"os"

	"GoLang_FRT_API_Tests/pkg/api"
	"GoLang_FRT_API_Tests/pkg/models"

	"github.com/stretchr/testify/assert"

)

const postIDEndpoint = "/posts/1"
const postEndpoint ="/posts"
const postNestedRoute="/posts/1/comments" 
const deserializationError = "❌ Error en la deserialización: %v"
const requestError = "❌ Error en la solicitud: %v"
const accessingEndpoint = "📡 Accediendo al Endpoint : %s"
const postIDLogFormat = "📝 ID del post: %d"
const postUserIDLogFormat = "📝 USER ID del post: %d"
const postTitleLogFormat = "📝 TITULO del post: %s"
const postBodyLogFormat = "📝 CONTENIDO del post: %s"

var logger *log.Logger

func TestMain(m *testing.M) {
	// Configurar el logger para escribir a un archivo
	logFile, err := os.OpenFile("../../reports/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("No se pudo crear el archivo de log:", err)
	}
	defer logFile.Close()

	// Configurar el logger para incluir timestamp
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	// Ejecutar los tests
	code := m.Run()
	os.Exit(code)
}

func testGetPosts(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Todos los Posts")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postEndpoint)

	resp, err := client.SendRequest(http.MethodGet, postEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var posts []models.Post
		err = json.Unmarshal(body, &posts)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, posts)
			logger.Printf("📊 Posts obtenidos: %v", len(posts))
			logger.Printf("✅ Test GET Todos los Posts completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testGetPostByID(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Post con ID=1")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postIDEndpoint)

	resp, err := client.SendRequest(http.MethodGet, postIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var post models.Post
		err = json.Unmarshal(body, &post)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, post)
			logger.Printf(postIDLogFormat, post.ID)
			logger.Printf(postUserIDLogFormat, post.UserID)
			logger.Printf(postTitleLogFormat, post.Title)
			logger.Printf(postBodyLogFormat, post.Body)
			logger.Printf("✅ Test GET Un Post por ID completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testGetCommentsByID(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Comentarios de un Post por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postNestedRoute)

	resp, err := client.SendRequest(http.MethodGet, postNestedRoute, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var comments []models.Comment
		err = json.Unmarshal(body, &comments)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, comments)
			logger.Printf("📊 Total de comentarios: %v", len(comments))
			for i := 0; i<(len(comments)); i++ {
				logger.Printf("📝 COMENTARIO %d del post: %s", comments[i].ID, comments[i].Body)
			}
			logger.Printf("✅ Test GET Comentarios de un Post por ID completado en %.2f", time.Since(startTime).Seconds())
		}
	}

}

func testPostNewPost(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando POST Creamos un nuevo Post")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postNestedRoute)

	newPost := models.Post {
		UserID: 1,
		Title: "Nuevo Post de Prueba",
		Body: "Cuerpo del nuevo Post",
	}
	
	resp, err := client.SendRequest(http.MethodPost, postEndpoint, newPost)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var createdPost models.Post
		err = json.Unmarshal(body, &createdPost)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.Equal(t,newPost.Title, createdPost.Title)
			logger.Printf("📊 NUEVO POST CREADO:")
			logger.Printf(postIDLogFormat, createdPost.ID)
			logger.Printf(postUserIDLogFormat, createdPost.UserID)
			logger.Printf(postTitleLogFormat, createdPost.Title)
			logger.Printf(postBodyLogFormat, createdPost.Body)
			logger.Printf("✅ Test POST Creamos un nuevo Post completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPutExistingPost(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PUT Actualizamos la totalidad de un Post")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postIDEndpoint)

	updatedPost := models.Post {
		UserID: 2,
		ID: 10,
		Title: "Post de Prueba Modificado",
		Body: "Cuerpo Modificado del Post",
	}
	
	resp, err := client.SendRequest(http.MethodPut, postIDEndpoint, updatedPost)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var updatedPutPost models.Post
		err = json.Unmarshal(body, &updatedPutPost)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			logger.Printf("📊 NUEVO POST MODIFICADO:")
			logger.Printf(postIDLogFormat, updatedPutPost.ID)
			logger.Printf(postUserIDLogFormat, updatedPutPost.UserID)
			logger.Printf(postTitleLogFormat, updatedPutPost.Title)
			logger.Printf(postBodyLogFormat, updatedPutPost.Body)
			logger.Printf("✅ Test PUT Actualizamos la totalidad de un Post completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPatchExistingPost(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PATCH Actualizamos una parte de un Post")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, postIDEndpoint)

	updatedPost := models.Post {
		Title: "Post de Prueba Modificado CON PATCH",
		Body: "Cuerpo Modificado del Post CON PATCH",
	}
	
	resp, err := client.SendRequest(http.MethodPatch, postIDEndpoint, updatedPost)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var updatedPatchPost models.Post
		err = json.Unmarshal(body, &updatedPatchPost)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			logger.Printf("📊 NUEVO POST MODIFICADO PARCIALMENTE:")
			logger.Printf(postIDLogFormat, updatedPatchPost.ID)
			logger.Printf(postUserIDLogFormat, updatedPatchPost.UserID)
			logger.Printf(postTitleLogFormat, updatedPatchPost.Title)
			logger.Printf(postBodyLogFormat, updatedPatchPost.Body)
			logger.Printf("✅ Test PATCH Actualizamos una parte de un Post completado en %.2f", time.Since(startTime).Seconds())
		}
	}

}

func testDeleteExistingPost(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando DELETE Borramos un Post")
	startTime := time.Now()

	logger.Printf(accessingEndpoint, postIDEndpoint)
	resp, err := client.SendRequest(http.MethodDelete, postIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		logger.Printf("📊 POST BORRADO OK")
		logger.Printf("✅ Test DELETE Borramos un Post completado en %.2f", time.Since(startTime).Seconds())
	}

}


func TestPostsEndpoints(t *testing.T) {
	client := api.NewClient()
	
	// GET Todos los Posts
	t.Run("GET /posts", func(t *testing.T) {
		testGetPosts(t, client)
	})
	// GET Un Post por ID
	t.Run("GET /posts/1", func(t *testing.T) {
		testGetPostByID(t, client)
	})
	

	// GET Comentarios de un Post por ID
	t.Run("GET /posts/1/comments", func(t *testing.T) {
		testGetCommentsByID(t, client)
	})

	// POST Creamos un nuevo Post
	t.Run("POST /posts", func(t *testing.T) {
		testPostNewPost(t, client)
	})

	// PUT Actualizamos la totalidad de un Post
	t.Run("PUT /posts/1", func(t *testing.T) {
		testPutExistingPost(t, client)

	})

	// PATCH Actualizamos una parte de un Post
	t.Run("PATCH /posts/1", func(t *testing.T) {
		testPatchExistingPost(t, client)
	})

	// DELETE Borramos un Post
	t.Run("DELETE /posts/1", func(t *testing.T) {
		testDeleteExistingPost(t, client)
	})


}
