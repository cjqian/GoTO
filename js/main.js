angular.module('app', ['ngReactGrid'])

.controller('InitCtrl', function($scope, $http, $log, ngReactGridCheckbox) {
    var ipAddress = "10.252.53.120";
    //initialization
    $scope.grid = {
        data: [],
        columnDefs: []
    }
    $scope.selections = [];
    getTableList();

    var checkboxGrid = new ngReactGridCheckbox($scope.selections, {
        batchToggle: true
    });
    //get list of tables
    function getTableList() {
        $http.get('http://' + ipAddress + ':8080/request/').then(function(resp) {
            $scope.tables = resp.data;
        }, function(err) {
            console.error('ERR', err);
        })
    }

    function setTable(data) {
		$scope.newRow = {}

        if (data.error != "") {
            alert(data.error);
        }

        checkboxGrid.setVisibleCheckboxState(false);
        $scope.editEnabled = false;

        //set grid
        $scope.grid = {
            data: data.response,
            columnDefs: data.colWrappers.concat(checkboxGrid),
            horizontalScroll: true
        }

        $scope.isTable = data.isTable;
        $scope.columns = data.columns;
    }

    $scope.clearCheckboxes = function() {
        checkboxGrid.setVisibleCheckboxState(false);
    }

    //GET
    $scope.get = function(table) {
        $http.get('http://' + ipAddress + ':8080/api/' + table).then(function(resp) {
            setTable(resp.data);
        }, function(err) {
            console.error('ERR', err);
            // err.status will contain the status code
        })
    }

    //GET
    $scope.update = function(table, parameters) {
        var tableName = angular.copy(table);

        if (typeof parameters !== 'undefined') {
            $http.get('http://' + ipAddress + ':8080/api/' + tableName + "?" + parameters).then(function(resp) {
                setTable(resp.data);
            }, function(err) {
                console.error('ERR', err);
            })
        } else {
            $scope.get(table);
        }
    }

    //DELETE
    $scope.delete = function(table, rows) {
        for (var i = 0; i < rows.length; i++) {
            $http.delete('http://' + ipAddress + ':8080/api/' + table + "/" + rows[i].id).then(function(resp) {
                setTable(resp.data);
            }, function(err) {
                console.error('ERR', err);
            })
        }
    }

    //DELETE
    $scope.deleteView = function(table) {
        $http.delete('http://' + ipAddress + ':8080/api/' + table).then(function(resp) {
            if (resp.data.error != "") {
                alert(resp.data.error);
            }

            location.reload();
            //make table
        }, function(err) {
            console.error('ERR', err);
            // err.status will contain the status code
        })

        getTableList();
    }


    //POST QUERY
    $scope.postView = function(newView) {
        var viewArray = new Array(newView);

        $http.post('http://' + ipAddress + ':8080/api/', viewArray).then(function(resp) {
            if (resp.data.error != "") {
                alert(resp.data.error);
            }
            location.reload();
        }, function(err) {
            console.error('ERR', err);
            // err.status will contain the status code
        })
    }

    $scope.post = function(table, row) {
        var rowArray = new Array(row);

        $http.post('http://' + ipAddress + ':8080/api/' + table, rowArray).then(function(resp) {
            setTable(resp.data);
        }, function(err) {
            console.error('ERR', err);
        })
    }

    //PUT
    $scope.put = function(table, row) {
        var rowArray = new Array(row);
        $http.put('http://' + ipAddress + ':8080/api/' + table + "/" + row.id, rowArray).then(function(resp) {
            setTable(resp.data);
        }, function(err) {
            console.error('ERR', err);
        })
    }
})
