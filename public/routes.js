angular.module('app.routes', [
  'ui.router','app.auth','app.list'
])
.config(function($stateProvider, $locationProvider) {
  'use strict';
  $locationProvider.html5Mode(true);
  $stateProvider
    .state('login', {
      url:'/',
      controller: 'LoginPageCtrl',
      templateUrl: '/views/login.html',
      requireLogin: false
      //title: 'Login',
    })
    .state('list', {
      url:'/list',
      controller: 'ListPageCtrl',
      templateUrl: '/views/list.html',
      requireLogin: true
      //title: 'List',
    });
})
