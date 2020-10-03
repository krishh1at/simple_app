$(function() {
    $('a[data-method]').on('click', function(e){
        var confirmMsg = $(this).attr('data-confirm');

        if(confirmMsg != '') {
            var condition = window.confirm(confirmMsg);

            if(condition == true){
                var url = $(this).attr('href');
                var dataMethod = $(this).attr('data-method');

                $.ajax(url, {
                    method: dataMethod
                });
            }
        }
    });

    $('form').on('submit', function(e){
        var url = $(this).attr('action');
        var dataMethod = $(this).attr('method');
        var formData = $(this).serialize();

        if(dataMethod.toLowerCase() == "post") {
            return
        }

        $.ajax(url, {
            method: dataMethod,
            data: formData
        });
    });
});
