angular.module('app.list',[])

.controller('ListPageCtrl', function($state,$scope, $http, $rootScope, $location) {
  if (!$rootScope.token) {
    $location.path('/');
    return;
  }

  $scope.ghUsername = '';
  $scope.items = [];
  $scope.total = 0;

  $scope.submit = function() {
    $http.put('/api/list', {token:$rootScope.token, ghUser:$scope.ghUsername})
    .then(function(resp) {
      for (i=0; i < resp.data.length; i++) {
        $scope.total += resp.data[i].stargazers_count;
      }
      $scope.items = resp.data;
      $scope.items.sort(function (a, b) {
        if (a.stargazers_count > b.stargazers_count) {
          return -1;
        }
        if (a.stargazers_count < b.stargazers_count) {
          return 1;
        }
        // a must be equal to b
        return 0;
      });
    });
  };

});
