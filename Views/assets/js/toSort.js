var selectAll = false;



/*
 * Au chargement de la vue
 */
$("#atrier").on("show", function () {
    $.ajax({
        url: "/tosort",
        type: "GET",
        success: function (data) {
            construct(data);
        }
    })
});
function construct(data) {
    $("tbody").empty();
    data = JSON.parse(data);
    for (x in data) {
        $("tbody").append('<tr><td><input type="checkbox" id="fichier' + x + '" /><label for="fichier' + x + '"></label></td><td>' + data[x] + '</td></tr>');
    }
}


/*
 * Checkbox selectAll dans header tableau
 */
$("#selectAll").on('click', function () {
    if (selectAll == true) {
        $('input[type=checkbox]').each(function () {
            $(this).prop("checked", false);
        });
        selectAll = false;
    } else {
        $('input[type=checkbox]').each(function () {
            $(this).prop("checked", true);
        });
        selectAll = true;
    }
});


/**
 * Bouton ajouter serie
 */
$("#addSerie").on('click', function () {
    // Vérif d'erreur a l'arrache
    var files = false;
    $('input[type=checkbox]').each(function () {
        if ($(this).prop('checked')) {
            files = true;
        }
    });

    if ($("#nomSerie").val() == "") {
        files = false;
    }


    if (files) {
        var nomSerie = $("#nomSerie").val();
        var data = {name: nomSerie};
        var episodes = [];
        $('input[type=checkbox]').each(function () {
            if ($(this).prop('checked')) {
                //Recup le nom de l'episode
                var episode = $(this).closest("td").next().text();
                //push ds tableau
                episodes.push(episode);
            }
        });
        data.episodes = episodes;
        $.ajax({
            url:"/topost",
            type: "post",
            data: "serie="+JSON.stringify(data),
            success: function(data) {
                console.log("post ok");
            },
            error: function(data) {
                console.log("error Post");
            }

        })
    }
    else {
        console.log("aucune serie");
    }
});
