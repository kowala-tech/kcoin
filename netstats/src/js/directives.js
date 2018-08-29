/* Directives */

angular
  .module("netStatsApp.directives", [])
  .directive("appVersion", [
    "version",
    function(version) {
      return function(scope, elm, attrs) {
        elm.text(version);
      };
    }
  ])
  .directive("sparkchart", function() {
    return {
      restrict: "E",
      scope: {
        data: "@"
      },
      compile: function(tElement, tAttrs, transclude) {
        tElement.replaceWith("<span>" + tAttrs.data + "</span>");

        return function(scope, element, attrs) {
          attrs.$observe("data", function(newValue) {
            element.html(newValue);
            element.sparkline("html", {
              defaultPixelsPerValue: 1,
              type: "line",
              spotColor: "#ffffff",
              highlightSpotColor: "#ffffff",
              maxSpotColor: false,
              minSpotColor: false,
              lineColor: "#ffffff",
              fillColor: false,
              width: "100%",
              tooltipSuffix: attrs.tooltipsuffix || "",
              tooltipPrefix: attrs.tooltipprefix || ""
            });
          });
        };
      }
    };
  })
