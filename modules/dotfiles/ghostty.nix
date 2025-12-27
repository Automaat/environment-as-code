{ config, ... }:

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

    # Everforest Light Med theme

    # Normal colors
    palette = 0=#6d9066
    palette = 1=#f06070
    palette = 2=#95d060
    palette = 3=#e8c060
    palette = 4=#60c8c0
    palette = 5=#e080c0
    palette = 6=#60d090
    palette = 7=#f5f0d0

    # Bright colors
    palette = 8=#95b890
    palette = 9=#ff4040
    palette = 10=#90b000
    palette = 11=#f0a000
    palette = 12=#2090d0
    palette = 13=#f050c0
    palette = 14=#20b870
    palette = 15=#ffffeb

    # Background and foreground
    background = #efebd4
    foreground = #456178

    # Cursor
    cursor-color = #f57d26

    # Selection
    selection-background = #eaedc8
    selection-foreground = #456178

    # ============================================================================
    # QUICK TERMINAL
    # ============================================================================

    keybind = global:ctrl+escape=toggle_quick_terminal
  '';
}
