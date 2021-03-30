//+build test_all integration

package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"github.com/thyago/directed-graph-go/config"
	"github.com/thyago/directed-graph-go/controllers"
)

type dgTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
}

func TestDGTestSuite(t *testing.T) {
	suite.Run(t, &dgTestSuite{})
}

func (s *dgTestSuite) SetupSuite() {
	server := controllers.NewServer()

	serverReady := make(chan bool)
	go server.Run()
	<-serverReady
}

func (s *dgTestSuite) TearDownSuite() {
	p, _ := os.FindProcess(syscall.Getpid())
	p.Signal(syscall.SIGINT)
}

// DIRECTED GRAPH

func (s *dgTestSuite) TestIntegration_CreateValidDirectedGraph() {
	reqStr := `{"name":"Graph 1"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%v", config.Config.ServerPort), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := ioutil.ReadAll(response.Body)
	s.NoError(err)

	s.Equal(`{"data":{"name":"Graph 1", "id": 1}}`, strings.Trim(string(byteBody), "\n"))
	response.Body.Close()
}

func TestIntegration_CreateInvalidDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_ListDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateValidDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateInvalidDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_GetExistingDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_GetNonExistingDirectedGraph(t *testing.T) {
	//TODO
}

func TestIntegration_DeleteDirectedGraph(t *testing.T) {
	//TODO
}

// NODE

func TestIntegration_CreateValidNode(t *testing.T) {
	//TODO
}

func TestIntegration_CreateInvalidNode(t *testing.T) {
	//TODO
}

func TestIntegration_ListNode(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateValidNode(t *testing.T) {
	//TODO
}

func TestIntegration_UpdateInvalidNode(t *testing.T) {
	//TODO
}

func TestIntegration_GetExistingNode(t *testing.T) {
	//TODO
}

func TestIntegration_GetNonExistingNode(t *testing.T) {
	//TODO
}

func TestIntegration_DeleteNode(t *testing.T) {
	//TODO
}

// EDGE

func TestIntegration_CreateValidEdge(t *testing.T) {
	//TODO
}

func TestIntegration_CreateInvalidEdge(t *testing.T) {
	//TODO
}

func TestIntegration_ListEdge(t *testing.T) {
	//TODO
}

func TestIntegration_GetExistingEdge(t *testing.T) {
	//TODO
}

func TestIntegration_GetNonExistingEdge(t *testing.T) {
	//TODO
}

func TestIntegration_DeleteEdge(t *testing.T) {
	//TODO
}
