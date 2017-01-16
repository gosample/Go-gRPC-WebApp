angular.module('app', [
  'ui.router','app.routes','app.auth','app.list'
])


.run(['$rootScope','$location','$state',function ($rootScope,$location,$state,$stateProvider) {

  $rootScope.$on('$stateChangeStart', function (event, toState, toParams) {

     data=localStorage.getItem('data');

     if(typeof data === 'string'){
       $rootScope.loginForm=false;
     }
     else{
       $rootScope.loginForm=true;
     }

    var requireLogin = toState.params.requireLogin;
    if (requireLogin && typeof data != 'string') {
      event.preventDefault();
      $state.go("login");
    }
  });

}])
