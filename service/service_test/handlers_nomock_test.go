package service

/*
func newCreateCarForm2() *url.Values {
	form := url.Values{}
	form.Set("make", "ford")
	form.Set("model", "mustang")
	return &form
}

func TestCreateCar2(t *testing.T) {
	// Given
	// var mockStore mock_store.Store = mock_store.Store{}
	var mockStore = new(mock_store.Store)

	testData := model.Car{
		Make:  "Citroen",
		Model: "c3",
	}

	// json := `{"car": {"make": "ford", "model": "mustang"}}`
	// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	// req, err := http.NewRequest("POST", "", r)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	mockStore.On("CreateCar", &testData).Return(nil)

	form := newCreateCarForm()
	req2, err := http.NewRequest("POST", "", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req2.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req2.Header.Add("Content-Length", strconv.Itoa(len(form.Encode())))

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(CreateCarHandler)
	hf.ServeHTTP(recorder, req2)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	mockStore.AssertExpectations(t)

}

func TestGetCarsHandler2(t *testing.T) {
	// Initialise the mock store
	var mockStore mock_store.Store = mock_store.Store{}
	// var mockClient mock_client.Client = mock_client.Client{}

	mockStore.On("GetCars").Return([]*model.Car{
		{
			Make:  "Citroen",
			Model: "c3",
		},
	}, nil).Once()

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	// var store12 Store
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(GetCarHandler)
	hf.ServeHTTP(recorder, req)
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := []model.Car{
		model.Car{Make: "citroen", Model: "c3"},
		// model.Car{"ford", "fiesta", 234},
	}
	c := []model.Car{}
	err = json.NewDecoder(recorder.Body).Decode(&c)
	if err != nil {
		t.Fatal(err)
	}
	actual := c[0]
	assert.Equal(t, expected, actual)

}
*/

// json := `{}`
// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
// response := &http.Response{Body: r, StatusCode: 200}
// mockClient.On("GET", "").Return(response)

// req, err := http.NewRequest("GET", "", nil)

// if err != nil {
// 	t.Fatal(err)
// }
// hf := http.HandlerFunc(aa.GetCarHandler)
// hf.ServeHTTP(recorder, req)

// recorder := httptest.NewRecorder()
// expected := model.Car{Make: "citroen", Model: "c3"}
// c := []model.Car{}
// json.NewDecoder(recorder.Body).Decode(&c)

// actual := c[0]
// if actual != expected {
// 	t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
// }
// mockStore.AssertExpectations(t)
