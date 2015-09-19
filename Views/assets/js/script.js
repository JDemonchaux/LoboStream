$(document).ready(function () {

    $("#formUpload").on('submit', function (e) {
        e.preventDefault();
        $("#formUpload").hide("slow");
        $(".loader").show("slow");

        //Pour envoyer plusieurs fichiers avec un FormData,
        //On parcourt tout nos fichiers qu'on append au formdata
        var formData = new FormData(this);
        $('#formUpload.files').each(function(i, file) {
            formData.append('video-'+i, file);
        });

        console.log(formData);
        $.ajax({
            url: "http://localhost:8080/upload",
            type: "POST",
            data: formData,
            //Les 3 options suivantes ont obligatoire pour le formData.
            cache: false,
            contentType: false,
            processData: false,
            xhr: function () {
                var request = $.ajaxSettings.xhr();
                if (request.upload) { // Test si la propriété upload existe.
                    request.addEventListener('progress', showProgress, false);
                }
                return request;
            },
            success: function (msg) {
                $(".loader").hide("slow");
                $("#formUpload").show("slow");
            },
            error: function (err) {
                $(".loader").hide("slow");
                $("#formUpload").show("slow");
                $('.progress').show("slow");
                $('.progress').html("<p>error : " + err + "</p>");
            }
        })
    });
});


function showProgress(e) {
    if (e.lengthComputable) {
        $('.progress').show("slow");
        $('.progress').html("<p>Progress : " + e.loaded + " sur total : " + e.total + "</p>");
    }
}