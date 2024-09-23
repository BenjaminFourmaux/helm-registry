package test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
	"gotest.tools/v3/assert"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var db *sql.DB

func TestStartup(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/test")
	if err != nil {
		t.Fatalf("Fail to request API: %v", err)
	}
	defer resp.Body.Close()

	assert.Assert(t, http.StatusOK == resp.StatusCode, "API start")
}

func TestInitState(t *testing.T) {
	solutionDir := filepath.Join("..")

	// Check if files exist
	files := []string{"/test/chart/test-nginx-1.0.0.tgz", "/test/chart/test-chart-0.1.0.tgz", "/backend/index.yaml", "backend/registry.db"}
	for _, file := range files {
		filePath := filepath.Join(solutionDir, file)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Fatalf("File %s does not exist", filePath)
		} else {
			assert.Assert(t, true, "File: "+file+" exist")
		}
	}

	// Testing SQL connexion
	dbConn, err := sql.Open("sqlite3", "../backend/registry.db")
	if err != nil {
		assert.Error(t, err, "Error SQL connection")
	} else {
		dbConn.Ping()
		db = dbConn
	}
}

func TestInitStateDB(t *testing.T) {
	assert.Assert(t, compareNbChartsInDB(2))
	t.Logf("Database has %d charts", 2)
}

func TestInitStateIndex(t *testing.T) {
	assert.Assert(t, compareNbChartsInIndex(2))
	t.Logf("Index.yaml has %d charts listed", 2)
}

func TestMoveChartOutside(t *testing.T) {
	t.Log("Move test-chart-0.1.0.tgz file outside CHART_DIR")
	assert.Assert(t, moveFile("chart/test-chart-0.1.0.tgz", "."))
	time.Sleep(2 * time.Second)
	t.Log("file: test-chart-0.1.0.tgz moved")
}

/*
TestRemoveChartStateIndex check if the previously removed chart is now not in index.yaml
*/
func TestRemoveChartStateIndex(t *testing.T) {
	assert.Assert(t, compareNbChartsInIndex(1))
	assert.Assert(t, getNameOfChartsInIndex()[0] == "test-nginx")
}

func TestRemoveChartStateDB(t *testing.T) {
	assert.Assert(t, compareNbChartsInDB(1))
	assert.Assert(t, getNameOfChartsInDB()[0] == "test-nginx")
}

func TestMoveChartInside(t *testing.T) {
	t.Log("Move test-chart-0.1.0.tgz file inside CHART_DIR")
	assert.Assert(t, moveFile("./test-chart-0.1.0.tgz", "./chart"))
	time.Sleep(2 * time.Second)
	t.Log("file: test-chart-0.1.0.tgz moved")
}

func TestAddChartStateIndex(t *testing.T) {
	assert.Assert(t, compareNbChartsInIndex(2))
	var chartsName = getNameOfChartsInIndex()
	assert.Assert(t, (chartsName[0] == "test-chart" && chartsName[1] == "test-nginx") || (chartsName[0] == "test-nginx" && chartsName[1] == "test-chart"))
}

func TestAddChartStateDB(t *testing.T) {
	assert.Assert(t, compareNbChartsInDB(2))
	assert.Assert(t, getNameOfChartsInDB()[0] == "test-nginx" && getNameOfChartsInDB()[1] == "test-chart")
}

//<editor-fold desc="Compare Functions">

func compareNbChartsInDB(assert int) bool {
	result, _ := db.Query(`SELECT * FROM charts`)

	var count = 0

	for result.Next() {
		count++
	}
	return assert == count
}

func compareNbChartsInIndex(assert int) bool {
	var indexFile Index
	file, err := os.ReadFile("../backend/index.yaml")
	if err != nil {

	} else {
		_ = yaml.Unmarshal(file, &indexFile)
		return len(indexFile.Entries) == assert
	}
	return false
}

func getNameOfChartsInIndex() []string {
	var list []string
	var indexFile Index

	file, err := os.ReadFile("../backend/index.yaml")
	if err != nil {

	} else {
		_ = yaml.Unmarshal(file, &indexFile)
		for _, elem := range indexFile.Entries {
			list = append(list, elem[0].Name)
		}
	}
	return list
}

func getNameOfChartsInDB() []string {
	var list []string
	result, _ := db.Query(`SELECT name FROM charts`)

	for result.Next() {
		var name string
		result.Scan(&name)

		list = append(list, name)
	}

	return list
}

func moveFile(src string, dest string) bool {
	err := os.Rename(src, filepath.Join(dest, filepath.Base(src)))
	if err != nil {
		println(err.Error())
	}
	return err == nil
}

//</editor-fold>
