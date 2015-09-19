$(document).ready(function () {

    $("#formUpload").on('submit', function (e) {
        e.preventDefault();
        $("#formUpload").hide("slow");
        $(".loader").show("slow");

        //Pour envoyer plusieurs fichiers avec un FormData,
        //On parcourt tout nos fichiers qu'on append au formdata

        var nbVideo = document.getElementById("uploadFiles").files.length;
        var current = 0;
        //for (var i = 0; i < ; i++) {
        //    var file = $("input[type=file]")[0].files[i];
        //    formData.append('videos', file);
        //}
        uploadVideo(current, nbVideo);

        //

    });


});

function uploadVideo(current, nbVideo) {
    var formData = new FormData();
    formData.append('videos', $("input[type=file]")[0].files[current]);
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
            var prog = parseInt( (current / nbVideo * 100), 10) + "%";
            $('.progressUpload').show("slow");
            $(".determinate").attr("style", "width: "+prog+";");
            current++;
            if (current <= nbVideo) {
                uploadVideo(current, nbVideo);
            } else {
                $(".loader").hide("slow");
                $("#formUpload").show("slow");
            }
        },
        error: function (err) {
            $(".loader").hide("slow");
            $("#formUpload").show("slow");
            $('.progressUpload').show("slow");
            $('.progressUpload').append("<p>error : " + err + "</p>");
        }
    })
}

function showProgress(e) {
    //if (e.lengthComputable) {
    //    console.log(e);
    //    $('.progressUpload').show("slow");
    //    $('.progressUpload').append("<p>Progress : " + parseInt( (e.loaded / e.total * 100), 10) + "%</p>");
    //} else {
    //    console.log("impossible de calculer le progress");
    //}
}