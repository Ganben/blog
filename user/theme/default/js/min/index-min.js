!function($){$("#admin-sign-in").on("click",function(){var i=$("#admin-login-panel");i.hasClass("hide")&&($(this).parent().find(".hide").removeClass("hide").animate({opacity:.6}),i.animate({opacity:1},250))}),$("#admin-login-bar").on("click",".admin-mask",function(){$(this).animate({opacity:0},200,function(){$(this).addClass("hide")}),$("#admin-login-panel").addClass("hide").css({opacity:0})}),$("#admin-mgr-btn").on("click",function(){$(this).parent().find(".hide").removeClass("hide").animate({opacity:.6}),$("#admin-mgr-panel").animate({left:0},200)}),$("#admin-mgr-bar").on("click",".admin-mask",function(){$(this).animate({opacity:0},200,function(){$(this).addClass("hide")}),$("#admin-mgr-panel").animate({left:-301},200)})}(window.jQuery),function($){$("#login-form").on("submit",function(){var i=$(this).serialize();return $.post("/api/user/login",i,function(i){if(console.log(i),i.meta.status){var n=i.data.token;$.cookie("token",n.token,{path:"/",expire:7}),window.location.reload()}}),!1})}(window.jQuery),function($){$("#admin-sign-out").on("click",function(){var i=$.cookie("token");$.post("/api/user/logout","token="+i,function(i){return i.meta.status?($.removeCookie("token"),void window.location.reload()):void $(this).show()}),$(this).hide()})}(window.jQuery);