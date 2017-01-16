angular.module('app.auth',[])
.controller('LoginPageCtrl', function($state,$scope, $http, $location, $rootScope) {
  $('body').css('background-image','url(../images/login.jpg)');
  $scope.user = {
    username: '',
    password: '',
  };

  data=localStorage.getItem('data');

  if(typeof data === 'string'){
    $scope.loginForm=false;
    $('listlink').css('visibility', 'visible');
    console.log($('listlink'));
  }
  else{
    $scope.loginForm=true;
    $('logoutlink').css('visibility', 'visible');
  }


  $scope.login = function() {
    $http.post('/api/login', $scope.user)
    .then(function(resp) {
      $rootScope.token = resp.data.token;
      localStorage.setItem('data',$rootScope.token);
      $location.path('/list');
    }, function(e) {
      alert ('error ', e);
    });
  };
})
