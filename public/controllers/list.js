angular.module('app.list',[])

.controller('ListPageCtrl', function($state,$scope, $http, $rootScope, $location) {

  $scope.ghUsername = '';
  $scope.items = [];
  $scope.total = 0;

  $('body').css('background-image','');

  $scope.submit = function() {
    data = localStorage.getItem('data');
    if( typeof data != 'string' ){
      $location.path('/');
    } else{
      $scope.total=0;
      $http.put('/api/list', {token:localStorage.getItem('data'), ghUser:$scope.ghUsername})
      .then(function(resp) {

        for (i=0; i < resp.data.list.length; i++) {
          if (typeof resp.data.list[i].stargazers_count != 'undefined')
          $scope.total += resp.data.list[i].stargazers_count;
        }
        $scope.items = resp.data.list;

        if(typeof $scope.items === 'object'){

        }

        $scope.items.sort(function (a, b) {
          if(typeof a.stargazers_count === 'undefined' && typeof b.stargazers_count === 'undefined'){
            return 0;
          }
          if(typeof a.stargazers_count === 'undefined'){
            return 1;
          }
          if(typeof b.stargazers_count === 'undefined'){
            return -1;
          }
          if (a.stargazers_count > b.stargazers_count) {
            return -1;
          }
          if (a.stargazers_count < b.stargazers_count) {
            return 1;
          }
          // a must be equal to b
          return 0;
        });
      }, function(e) {
        $scope.loginError="Wrong Username / Password. Please retry."
      });
    }
    };

});
