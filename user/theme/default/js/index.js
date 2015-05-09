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
(function ($) {
    $('#login-form').on("submit", function () {
        var data = $(this).serialize();
        $.post("/api/user/login", data, function (result) {
            console.log(result);
            if (result.meta.status) {
                var token = result.data.token;
                $.cookie("token", token.token, {
                    path: "/",
                    expire: 7
                });
                window.location.reload();
            } else {

            }
        });
        return false;
    });
})(window.jQuery);

// logout logic
(function ($) {
    $('#admin-sign-out').on("click", function () {
        var token = $.cookie("token");
        $.post("/api/user/logout", "token=" + token, function (result) {
            if (result.meta.status) {
                $.removeCookie("token");
                window.location.reload();
                return;
            }
            $(this).show();
        });
        $(this).hide();
    });
})(window.jQuery);