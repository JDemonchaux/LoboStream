(function ($) {
    $.each(['show', 'hide'], function (i, ev) {
        var el = $.fn[ev];
        $.fn[ev] = function () {
            this.trigger(ev);
            return el.apply(this, arguments);
        };
    });
})(jQuery);

$(document).ready(function () {
    $(document).on('click', ".btClear" ,function () {
        $(this).parent().parent().remove();
    });


});