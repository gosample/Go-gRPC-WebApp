angular.module('app.list',[])

.controller('ListPageCtrl', function($state,$scope, $http, $rootScope, $location) {

  $scope.ghUsername = '';
  $scope.items = [];
  $scope.total = 0;
  $scope.githubUser="";

  $('body').css('background-image','');

  $scope.submit = function() {
    $scope.items =[];
    $scope.total=0;
    $scope.listError="";
    $scope.githubUser=$scope.ghUsername;

    data = Cookies.get('data');
    if( typeof data != 'string' ){
      $location.path('/');
    } else{
      $http.put('/api/list', {token:data, ghUser:$scope.ghUsername})
      .then(function(resp) {
        if(resp.data.token == "")
        {
          $location.path('/');
        }
        if(typeof resp.data.list === 'undefined'){
          $scope.listError="Github user not found! :("
        }
        else{
          for (i=0; i < resp.data.list.length; i++) {
            if (typeof resp.data.list[i].stargazers_count != 'undefined')
            $scope.total += resp.data.list[i].stargazers_count;
          }
          $scope.items = resp.data.list;
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
        }
      }, function(e) {
        $scope.listError="We're sorry! Please try again!"
      });
    }
    };

});
