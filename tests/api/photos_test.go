// tests/api/photos_test.go

package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
	

	"GoLang_FRT_API_Tests/pkg/api"
	"GoLang_FRT_API_Tests/pkg/models"

	"github.com/stretchr/testify/assert"

)

const photoIDEndpoint = "/photos/1"
const photoEndpoint ="/photos"
const photoByAlbumIDEndpoint="/albums/1/photos" 

const photoIDLogFormat = "ℹ️ ID de la FOTO: %d"
const albumIDLogFormat = "ℹ️ ID del ALBUM: %d"
const photoTitleLogFormat = "📷 TITULO de la FOTO: %s"
const photoURLLogFormat = "🌏 URL de la FOTO: %s"
const photoThumbnailURLLogFormat = "🌏 URL de la miniatura de la FOTO: %s"


func testGetPhotos(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Todas las Fotos")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoEndpoint)

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

		var photos []models.Photo
		err = json.Unmarshal(body, &photos)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photos)
			logger.Printf("📊 Total de fotos: %v", len(photos))
			logger.Printf("✅ Test GET Todas las FOTOS completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testGetPhotoByID(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Photo con ID=1")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoIDEndpoint)

	resp, err := client.SendRequest(http.MethodGet, photoIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var photo models.Photo
		err = json.Unmarshal(body, &photo)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photo)
			logger.Printf(photoIDLogFormat, photo.ID)
			logger.Printf(albumIDLogFormat, photo.AlbumID)
			logger.Printf(photoTitleLogFormat, photo.Title)
			logger.Printf(photoURLLogFormat, photo.URL)
			logger.Printf(photoThumbnailURLLogFormat, photo.ThumbnailURL)
			logger.Printf("✅ Test GET Un Photo por ID completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testGetPhotosByAlbumID(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Fotos de un Album por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoByAlbumIDEndpoint)

	resp, err := client.SendRequest(http.MethodGet, photoByAlbumIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var photos []models.Photo
		err = json.Unmarshal(body, &photos)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photos)
			logger.Printf("📊 Total de fotos: %v", len(photos))
			for i := 0; i<(len(photos)); i++ {
				logger.Printf("📷 FOTO %d del album: %s", photos[i].ID, photos[i].Title)
			}
			logger.Printf("✅ Test GET Fotos de un Album por ID completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPostNewPhoto(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando POST Creamos una nueva Foto")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoEndpoint)

	newPhoto := models.Photo {
		AlbumID: 1,
		ID: 5001,
		Title: "Nueva Foto de Prueba",
		URL: "https://via.placeholder.com/600/92c952",
		ThumbnailURL: "https://via.placeholder.com/150/92c952",
	}

	resp, err := client.SendRequest(http.MethodPost, photoEndpoint, newPhoto)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var photo models.Photo
		err = json.Unmarshal(body, &photo)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photo)
			logger.Printf(photoIDLogFormat, photo.ID)
			logger.Printf(albumIDLogFormat, photo.AlbumID)
			logger.Printf(photoTitleLogFormat, photo.Title)
			logger.Printf(photoURLLogFormat, photo.URL)
			logger.Printf(photoThumbnailURLLogFormat, photo.ThumbnailURL)
			logger.Printf("✅ Test POST Creamos una nueva Foto completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPutExistingPhoto(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PUT Actualizamos una Foto")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoIDEndpoint)

	newPhoto := models.Photo {
		AlbumID: 1,
		ID: 1,
		Title: "Foto de Prueba Modificada",
		URL: "https://www.urlmodificada.com",
		ThumbnailURL: "https://www.thumbnailurlmodificada.com",
	}

	resp, err := client.SendRequest(http.MethodPut, photoIDEndpoint, newPhoto)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var photo models.Photo
		err = json.Unmarshal(body, &photo)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photo)
			logger.Printf(photoIDLogFormat, photo.ID)
			logger.Printf(albumIDLogFormat, photo.AlbumID)
			logger.Printf(photoTitleLogFormat, photo.Title)
			logger.Printf(photoURLLogFormat, photo.URL)
			logger.Printf(photoThumbnailURLLogFormat, photo.ThumbnailURL)
			logger.Printf("✅ Test PUT Actualizamos una Foto completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPatchExistingPhoto(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PATCH Actualizamos una Foto")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoIDEndpoint)

	updatedPhoto := models.Photo {
		Title: "Foto de Prueba Modificada CON PATCH",
		URL: "https://www.urlmodiconPatch.com",
		ThumbnailURL: "https://www.thumbnailurlmodiconPatch.com",
	}

	resp, err := client.SendRequest(http.MethodPatch, photoIDEndpoint, updatedPhoto)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err:= io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var photo models.Photo
		err = json.Unmarshal(body, &photo)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		}else{
			assert.NoError(t, err)
			assert.NotEmpty(t, photo)
			logger.Printf(photoIDLogFormat, photo.ID)
			logger.Printf(albumIDLogFormat, photo.AlbumID)
			logger.Printf(photoTitleLogFormat, photo.Title)
			logger.Printf(photoURLLogFormat, photo.URL)
			logger.Printf(photoThumbnailURLLogFormat, photo.ThumbnailURL)
			logger.Printf("✅ Test PATCH Actualizamos una Foto completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testDeleteExistingPhoto(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando DELETE Eliminamos una Foto")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, photoIDEndpoint)

	resp, err := client.SendRequest(http.MethodDelete, photoIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	}else{
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		logger.Printf("✅ Test DELETE Eliminamos una Foto completado en %.2f", time.Since(startTime).Seconds())
	}
}


func TestPhotosEndpoints(t *testing.T) {
	client := api.NewClient()
	
	// GET Todas las fotos
	t.Run("GET /photos", func(t *testing.T) {
		testGetPhotos(t, client)
	})
	// GET Una Foto por ID
	t.Run("GET /photos/1", func(t *testing.T) {
		testGetPhotoByID(t, client)
	})
	
	// GET Fotos de un Album por ID
	t.Run("GET /albums/1/photos", func(t *testing.T) {
		testGetPhotosByAlbumID(t, client)
	})

	// POST Creamos una nueva Foto
	t.Run("POST /photos", func(t *testing.T) {
		testPostNewPhoto(t, client)
	})

	// PUT Actualizamos la totalidad de una Foto
	t.Run("PUT /photos/1", func(t *testing.T) {
		testPutExistingPhoto(t, client)
	})

	// PATCH Actualizamos una parte de una Foto
	t.Run("PATCH /photos/1", func(t *testing.T) {
		testPatchExistingPhoto(t, client)
	})

	// DELETE Borramos una Foto
	t.Run("DELETE /photos/1", func(t *testing.T) {
		testDeleteExistingPhoto(t, client)
	})
}
