// admin bar logic
(function ($) {
    $('#admin-sign-in').on("click", function () {
        var $adminLoginPanel = $("#admin-login-panel");
        if ($adminLoginPanel.hasClass("hide")) {
            $(this).parent().find(".hide").removeClass("hide").animate({opacity: 0.6});
            $adminLoginPanel.animate({opacity: 1}, 250);
        }
    });
    $('#admin-login-bar').on("click", ".admin-mask", function () {
        $(this).animate({opacity: 0}, 200, function () {
            $(this).addClass("hide")
        });
        $("#admin-login-panel").addClass("hide").css({opacity: 0});
    });
    $('#admin-mgr-btn').on("click", function () {
        $(this).parent().find(".hide").removeClass("hide").animate({opacity: 0.6});
        $('#admin-mgr-panel').animate({left: 0}, 200);
    });
    $('#admin-mgr-bar').on("click", ".admin-mask", function () {
        $(this).animate({opacity: 0}, 200, function () {
            $(this).addClass("hide")
        });
        $('#admin-mgr-panel').animate({left: -301}, 200);
    })
})(window.jQuery);

// login logic