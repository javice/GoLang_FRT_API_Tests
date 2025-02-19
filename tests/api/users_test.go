// tests/api/users_test.go

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

const userIDEndpoint = "/users/1"
const userEndpoint = "/users"
const userNestedRouteAlbums = "/users/1/albums"
const userNestedRouteTodos = "/users/1/todos"
const userNestedRoutePosts = "/users/1/posts"

const userID = "📊 Usuario Encontrado: %+v"
const userName = "🙋🏻‍♂️ Nombre: %+v"
const userUsername = "✍🏼 User name: %v"
const userMail = "📨 Email: %+v"
const userAddress = "🏠 Dirección completa: %+v %v %v %v %v %v "
const userPhone = "📞 Teléfono: %+v"
const userWebsite = "🌐 Web: %+v"
const userCompany = "🏢 Empresa: %+v %v %v"


func testGetUsers(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Todos los Usuarios")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userEndpoint)

	resp, err := client.SendRequest(http.MethodGet, userEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var users []models.User
		err = json.Unmarshal(body, &users)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, users)
			logger.Printf("📊 Usuarios totales: %v", len(users))
			logger.Printf("✅ Test GET Todos los Usuarios completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}
func testGetUserById(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Un Usuario por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userIDEndpoint)

	resp, err := client.SendRequest(http.MethodGet, userIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, user)
			logger.Printf(userID, user.ID)
			logger.Printf(userName, user.Name)
			logger.Printf(userMail, user.Email)
			logger.Printf(userUsername, user.Username)
			logger.Printf(userAddress, user.Address.Street, user.Address.Suite, user.Address.City, user.Address.Zipcode, user.Address.Geo.Lat, user.Address.Geo.Lng)
			logger.Printf(userPhone, user.Phone)
			logger.Printf(userWebsite, user.Website)
			logger.Printf(userCompany, user.Company.Name, user.Company.CatchPhrase, user.Company.Bs)
			logger.Printf("✅ Test GET Usuario por ID completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testGetUserAlbumsById(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Albums de un Usuario por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userNestedRouteAlbums)

	resp, err := client.SendRequest(http.MethodGet, userNestedRouteAlbums, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var albums []models.Album
		err = json.Unmarshal(body, &albums)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, albums)
			logger.Printf("📊 Albums totales: %v", len(albums))
			logger.Printf("✅ Test GET Albums de un Usuario por ID completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testTodosByUserId(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Todos de un Usuario por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userNestedRouteTodos)

	resp, err := client.SendRequest(http.MethodGet, userNestedRouteTodos, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var todos []models.Todo
		err = json.Unmarshal(body, &todos)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, todos)
			logger.Printf("📊 Todos totales: %v", len(todos))
			logger.Printf("✅ Test GET Todos de un Usuario por ID completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testPostsByUserId(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando GET Posts de un Usuario por ID")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userNestedRoutePosts)

	resp, err := client.SendRequest(http.MethodGet, userNestedRoutePosts, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var posts []models.Post
		err = json.Unmarshal(body, &posts)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, posts)
			logger.Printf("🙋🏻‍♂️ Usuario: %v", posts[0].UserID)
			logger.Printf("📊 Posts totales: %v", len(posts))
			logger.Printf("✅ Test GET Posts de un Usuario por ID completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testCreateUser(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando POST Crear Usuario")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userEndpoint)

	newUser := models.User{
		ID:       101,
		Name:     "Pepito Prueba",
		Username: "pepitoPruebaTest",
		Email:    "pepito@prueba.com",
		Address: struct {
			Street  string `json:"street"`
			Suite   string `json:"suite"`
			City    string `json:"city"`
			Zipcode string `json:"zipcode"`
			Geo     struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"geo"`
		}{
			Street:  "Paseo de Los Melancólicos 10",
			Suite:   "2-DCHA",
			City:    "Madrid",
			Zipcode: "28010",
			Geo: struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			}{
				Lat: "12345",
				Lng: "67890",
			},
		},
		Phone:   "913233480",
		Website: "www.pepito.com",
		Company: struct {
			Name        string `json:"name"`
			CatchPhrase string `json:"catchPhrase"`
			Bs          string `json:"bs"`
		}{
			Name:        "Pepito Corp",
			CatchPhrase: "Pepe Pepito",
			Bs:          "Bs de Pepito",
		},
	}

	resp, err := client.SendRequest(http.MethodPost, userEndpoint, newUser)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, user)
			logger.Printf("📊 DATOS DEL NUEVO USUARIO:")
			logger.Printf(userID, user.ID)
			logger.Printf(userName, user.Name)
			logger.Printf(userMail, user.Email)
			logger.Printf(userUsername, user.Username)
			logger.Printf(userAddress, user.Address.Street, user.Address.Suite, user.Address.City, user.Address.Zipcode, user.Address.Geo.Lat, user.Address.Geo.Lng)
			logger.Printf(userPhone, user.Phone)
			logger.Printf(userWebsite, user.Website)
			logger.Printf(userCompany, user.Company.Name, user.Company.CatchPhrase, user.Company.Bs)
			logger.Printf("✅ Test POST Crear Usuario completado en %.2f", time.Since(startTime).Seconds())

		}
	}
}

func testPutExistingUser(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PUT Actualizar Usuario")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userEndpoint)

	updatedUser := models.User{
		Name:     "Pepito Modificado",
		Username: "pepitoPruebaTestModificado",
		Email:    "pepito.modificado@prueba.com",
		Address: struct {
			Street  string `json:"street"`
			Suite   string `json:"suite"`
			City    string `json:"city"`
			Zipcode string `json:"zipcode"`
			Geo     struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			} `json:"geo"`
		}{
			Street:  "Calle San Epifanio 5 ",
			Suite:   "2-DCHA",
			City:    "Madrid",
			Zipcode: "28009",
			Geo: struct {
				Lat string `json:"lat"`
				Lng string `json:"lng"`
			}{
				Lat: "11111",
				Lng: "22222",
			},
		},
		Phone:   "913233481",
		Website: "www.pepito.modificado.com",
		Company: struct {
			Name        string `json:"name"`
			CatchPhrase string `json:"catchPhrase"`
			Bs          string `json:"bs"`
		}{
			Name:        "Pepito Corp V2",
			CatchPhrase: "Pepe Pepito Pepe",
			Bs:          "Bs de Pepito Modificado",
		},
	}

	resp, err := client.SendRequest(http.MethodPut, userIDEndpoint, updatedUser)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)	
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, user)
			logger.Printf("📊 DATOS DEL USUARIO MODIFICADO:")
			logger.Printf(userName, user.Name)
			logger.Printf(userMail, user.Email)
			logger.Printf(userUsername, user.Username)
			logger.Printf(userAddress, user.Address.Street, user.Address.Suite, user.Address.City, user.Address.Zipcode, user.Address.Geo.Lat, user.Address.Geo.Lng)
			logger.Printf(userPhone, user.Phone)
			logger.Printf(userWebsite, user.Website)
			logger.Printf(userCompany, user.Company.Name, user.Company.CatchPhrase, user.Company.Bs)
			logger.Printf("✅ Test PUT Actualizar Usuario completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testPatchExistingUser(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando PATCH Actualizar Usuario")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userEndpoint)

	updatedUser := models.User{
		Name:     "Pepito Modificado Patch",
		Username: "pepitoPruebaTestModificadoPatch",
		Email:    "pepito.modificado.patch@prueba.com",
	}

	resp, err := client.SendRequest(http.MethodPatch, userIDEndpoint, updatedUser)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		assert.NoError(t, err)
		defer resp.Body.Close()

		var user models.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			logger.Printf(deserializationError, err)
			t.Fatal(err)
		} else {
			assert.NoError(t, err)
			assert.NotEmpty(t, user)
			logger.Printf("📊 DATOS DEL USUARIO MODIFICADO PARCIALMENTE CON METODO PATCH:")
			logger.Printf(userName, user.Name)
			logger.Printf(userMail, user.Email)
			logger.Printf(userUsername, user.Username)
			logger.Printf("✅ Test PATCH Actualizar Usuario completado en %.2f", time.Since(startTime).Seconds())
		}
	}
}

func testDeleteExistingUser(t *testing.T, client *api.Client) {
	logger.Printf("🚀 Iniciando DELETE Borrar Usuario")
	startTime := time.Now()
	logger.Printf(accessingEndpoint, userIDEndpoint)

	resp, err := client.SendRequest(http.MethodDelete, userIDEndpoint, nil)
	if err != nil {
		logger.Printf(requestError, err)
		t.Fatal(err)
	} else {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		logger.Printf("✅ Test DELETE Borrar Usuario completado en %.2f", time.Since(startTime).Seconds())
	}
}

func TestUsersEndpoints(t *testing.T) {
	client := api.NewClient()

	// GET Todos los Usuarios
	t.Run("GET /users", func(t *testing.T) {
		testGetUsers(t, client)
	})

	// GET Un Usuario por ID
	t.Run("GET /user/1", func(t *testing.T) {
		testGetUserById(t, client)
	})

	// GET Albums de un Usuario por ID
	t.Run("GET /users/1/albums", func(t *testing.T) {
		testGetUserAlbumsById(t, client)
	})
		
	// GET Todos de un Usuario por ID
	t.Run("GET /users/1/todos", func(t *testing.T) {
		testTodosByUserId(t, client)
	})

	// GET Posts de un Usuario por ID
	t.Run("GET /users/1/posts", func(t *testing.T) {
		testPostsByUserId(t, client)
	})

	// POST Creamos un nuevo Usuario
	t.Run("POST /users", func(t *testing.T) {
		testCreateUser(t, client)
	})

	// PUT Actualizamos la totalidad de un Usuario
	t.Run("PUT /users/1", func(t *testing.T) {
		testPutExistingUser(t, client)
	})

	// PATCH Actualizamos una parte de un Usuario
	t.Run("PATCH /users/1", func(t *testing.T) {
		testPatchExistingUser(t, client)
	})

	// DELETE Borramos un Usuario
	t.Run("DELETE /users/1", func(t *testing.T) {
		testDeleteExistingUser(t, client)
	})

}
