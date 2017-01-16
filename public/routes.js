angular.module('app.routes', [
  'ui.router','app.auth','app.list'
])
.config(function($stateProvider, $locationProvider, $urlRouterProvider) {
  'use strict';
  $locationProvider.html5Mode(true);
  $stateProvider
    .state('login', {
      url:'/',
      controller: 'LoginPageCtrl',
      templateUrl: '/views/login.html',
      params: {requireLogin:false}
    })
    .state('list', {
      url:'/list',
      controller: 'ListPageCtrl',
      templateUrl: '/views/list.html',
      params: {requireLogin: true}
    })
    .state('logout', {
      url:'/logout',
      controller: 'LogoutCtrl',
      templateUrl: '/views/login.html',
      params: {requireLogin: true}
    })

    $urlRouterProvider.otherwise('/');
})
