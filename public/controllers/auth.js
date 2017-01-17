angular.module('app.auth',[])
.controller('LoginPageCtrl', function($state,$scope, $http, $location, $rootScope) {
  $('body').css('background-image','url(../images/login.jpg)');
  $scope.user = {
    username: '',
    password: '',
  };

  data= Cookies.get('data');

  if(typeof data === 'string'){
    $scope.loginForm=false;
  }
  else{
    $scope.loginForm=true;
  }


  $scope.login = function() {
    $http.post('/api/login', $scope.user)
    .then(function(resp) {
      $rootScope.token = resp.data.token;
      if(typeof $rootScope.token === 'undefined'){
          $scope.error="Wrong Username / Password. Please retry."
      }else{
          Cookies.set('data',$rootScope.token);
          $location.path('/list');
      }
    }, function(e) {
      $scope.error="Server Error! Please try again later."
    });
  };
})

.controller('LogoutCtrl', function($state,$scope, $http, $location, $rootScope,$window) {
  $('body').css('background-image','url(../images/login.jpg)');
  Cookies.remove('data');
  $location.path('/');
  $window.location.reload();
})

.controller('CreateUserCtrl', function($state,$scope, $http, $location, $rootScope,$window) {
  $('body').css('background-image','url(../images/login.jpg)');

  $scope.newuser = {
    username: '',
    password: '',
  };

  $scope.createUser = function() {
    $http.post('/api/createuser', $scope.newuser)
    .then(function(resp) {
      $rootScope.token = resp.data.token;
      if(typeof $rootScope.token === 'undefined'){
          $scope.nuerror="User already exists. "
      }else{
          Cookies.set('data',$rootScope.token);
          $location.path('/list');
      }
    }, function(e) {
      $scope.nuerror="Server Error! Please try again later."
    });
  };

})
