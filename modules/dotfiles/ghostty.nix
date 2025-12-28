{ config, ... }:

let
  # Theme selection - change this to switch themes
  theme = "nord-light";

  # Theme definitions
  themes = {
    nord-light = {
      name = "Nord Light";
      background = "#eceff4";
      foreground = "#2e3440";
      cursor-color = "#5e81ac";
      cursor-text = "#eceff4";
      selection-background = "#d8dee9";
      selection-foreground = "#2e3440";
      palette = {
        "0" = "#3b4252";
        "1" = "#bf616a";
        "2" = "#a3be8c";
        "3" = "#ebcb8b";
        "4" = "#81a1c1";
        "5" = "#b48ead";
        "6" = "#88c0d0";
        "7" = "#e5e9f0";
        "8" = "#4c566a";
        "9" = "#bf616a";
        "10" = "#a3be8c";
        "11" = "#ebcb8b";
        "12" = "#81a1c1";
        "13" = "#b48ead";
        "14" = "#8fbcbb";
        "15" = "#eceff4";
      };
    };

    nord-dark = {
      name = "Nord Dark";
      background = "#2e3440";
      foreground = "#d8dee9";
      cursor-color = "#81a1c1";
      cursor-text = "#2e3440";
      selection-background = "#4c566a";
      selection-foreground = "#d8dee9";
      palette = {
        "0" = "#3b4252";
        "1" = "#bf616a";
        "2" = "#a3be8c";
        "3" = "#ebcb8b";
        "4" = "#81a1c1";
        "5" = "#b48ead";
        "6" = "#88c0d0";
        "7" = "#e5e9f0";
        "8" = "#4c566a";
        "9" = "#bf616a";
        "10" = "#a3be8c";
        "11" = "#ebcb8b";
        "12" = "#81a1c1";
        "13" = "#b48ead";
        "14" = "#8fbcbb";
        "15" = "#eceff4";
      };
    };

    everforest-light = {
      name = "Everforest Light";
      background = "#fdf6e3";
      foreground = "#5c6a72";
      cursor-color = "#f57d26";
      cursor-text = "#fdf6e3";
      selection-background = "#f2efdf";
      selection-foreground = "#5c6a72";
      palette = {
        "0" = "#5c6a72";
        "1" = "#f85552";
        "2" = "#8da101";
        "3" = "#dfa000";
        "4" = "#3a94c5";
        "5" = "#df69ba";
        "6" = "#35a77c";
        "7" = "#dfddc8";
        "8" = "#859289";
        "9" = "#f85552";
        "10" = "#8da101";
        "11" = "#dfa000";
        "12" = "#3a94c5";
        "13" = "#df69ba";
        "14" = "#35a77c";
        "15" = "#fffbef";
      };
    };
  };

  selectedTheme = themes.${theme};
in
{
  xdg.configFile."ghostty/config".text = ''
    # Font
    font-family = Iosevka
    font-family = Symbols Nerd Font Mono
    font-size = 23

    # Window
    window-width = 240
    window-height = 80

    # ============================================================================
    # PERFORMANCE - Critical for Claude Code (heavy output)
    # ============================================================================

    # Large scrollback for long AI responses
    scrollback-limit = 200000000

    # Smooth scrolling during code generation
    cursor-style-blink = false

    # ============================================================================
    # MOUSE
    # ============================================================================

    mouse-hide-while-typing = true
    mouse-scroll-multiplier = 3

    # ============================================================================
    # KEYBOARD
    # ============================================================================

    keybind = shift+enter=text:\n

    # ============================================================================
    # THEME: ${selectedTheme.name}
    # ============================================================================

    # Normal colors
    palette = 0=${selectedTheme.palette."0"}
    palette = 1=${selectedTheme.palette."1"}
    palette = 2=${selectedTheme.palette."2"}
    palette = 3=${selectedTheme.palette."3"}
    palette = 4=${selectedTheme.palette."4"}
    palette = 5=${selectedTheme.palette."5"}
    palette = 6=${selectedTheme.palette."6"}
    palette = 7=${selectedTheme.palette."7"}

    # Bright colors
    palette = 8=${selectedTheme.palette."8"}
    palette = 9=${selectedTheme.palette."9"}
    palette = 10=${selectedTheme.palette."10"}
    palette = 11=${selectedTheme.palette."11"}
    palette = 12=${selectedTheme.palette."12"}
    palette = 13=${selectedTheme.palette."13"}
    palette = 14=${selectedTheme.palette."14"}
    palette = 15=${selectedTheme.palette."15"}

    # Background and foreground
    background = ${selectedTheme.background}
    foreground = ${selectedTheme.foreground}

    # Cursor
    cursor-color = ${selectedTheme.cursor-color}

    # Selection
    selection-background = ${selectedTheme.selection-background}
    selection-foreground = ${selectedTheme.selection-foreground}

    # ============================================================================
    # QUICK TERMINAL
    # ============================================================================

    keybind = global:ctrl+escape=toggle_quick_terminal
  '';
}
