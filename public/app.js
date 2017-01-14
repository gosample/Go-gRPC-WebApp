angular.module('app', [
  'ui.router','app.routes','app.auth','app.list'
])


.run(['$rootScope','$location','$state',function ($rootScope,$location,$stateProvider) {

  $rootScope.$on('$stateChangeStart', function (event, toState, toParams) {
    var requireLogin = $stateProvider.requireLogin;

    if (requireLogin && typeof $rootScope.currentUser === 'undefined') {
      event.preventDefault();
      $location.path('/');
    }
  });

}])
