angular.module('app.auth',[])
.controller('LoginPageCtrl', function($state,$scope, $http, $location, $rootScope) {
  $scope.user = {
    username: '',
    password: '',
  };

  $scope.login = function() {
    $http.post('/api/login', $scope.user)
    .then(function(resp) {
      $rootScope.token = resp.data.token;
      $location.path('/list');
    }, function(e) {
      alert ('error ', e);
    });
  };
})
