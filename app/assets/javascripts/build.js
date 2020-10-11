$(function() {
    $('a[data-method]').on('click', function(e){
        var confirmMsg = $(this).attr('data-confirm');

        if(confirmMsg != undefined) {
          var condition = window.confirm(confirmMsg);

          if(condition == true){
            deleteAjaxRequest(this);
          }
        } else {
          deleteAjaxRequest(this);
        }
    });

    $('form').on('submit', function(e){
        var url = $(this).attr('action');
        var dataMethod = $(this).attr('method');
        var formData = $(this).serialize();

        if(dataMethod.toLowerCase() != "post") {
            $.ajax(url, {
                method: dataMethod,
                data: formData
            });
        }
    });
});

function deleteAjaxRequest(node) {
  var url = $(node).attr('data-href');
  var dataMethod = $(node).attr('data-method');
  console.log(dataMethod);

  $.ajax(url, {
    method: dataMethod
  });
}
