angular.module('app.list',[])

.controller('ListPageCtrl', function($state,$scope, $http, $rootScope, $location) {

  $scope.ghUsername = '';
  $scope.items = [];
  $scope.total = 0;

  $('body').css('background-image','');

  $scope.submit = function() {
    console.log("in in");
    $http.put('/api/list', {token:$rootScope.token, ghUser:$scope.ghUsername})
    .then(function(resp) {
      console.log(resp.data.list.length)
      console.log(resp.data);
      for (i=0; i < resp.data.list.length; i++) {
        if (typeof resp.data.list[i].stargazers_count != 'undefined')
        $scope.total += resp.data.list[i].stargazers_count;
      }
      $scope.items = resp.data;
      if(typeof $scope.items === 'object'){
        console.log(typeof $scope.items)
        console.log($scope.total)
      }

      // $scope.items.sort(function (a, b) {
      //   if (a.stargazers_count > b.stargazers_count) {
      //     return -1;
      //   }
      //   if (a.stargazers_count < b.stargazers_count) {
      //     return 1;
      //   }
      //   // a must be equal to b
      //   return 0;
      // });
    });
  };

});
