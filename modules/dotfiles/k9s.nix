{ config, ... }:

{
  xdg.configFile."k9s/config.yaml".text = ''
    k9s:
      liveViewAutoRefresh: false
      screenDumpDir: ${config.home.homeDirectory}/Library/Application Support/k9s/screen-dumps
      refreshRate: 2
      maxConnRetry: 5
      readOnly: false
      noExitOnCtrlC: false
      ui:
        enableMouse: false
        headless: false
        logoless: false
        crumbsless: false
        noIcons: false
        skin: nord-light
      skipLatestRevCheck: false
      disablePodCounting: false
      shellPod:
        image: busybox:1.35.0
        namespace: default
        limits:
          cpu: 100m
          memory: 100Mi
      imageScans:
        enable: false
        exclusions:
          namespaces: []
          labels: {}
      logger:
        tail: 100
        buffer: 5000
        sinceSeconds: -1
        fullScreenLogs: false
        textWrap: false
        showTime: false
      thresholds:
        cpu:
          critical: 90
          warn: 70
        memory:
          critical: 90
          warn: 70
  '';

  xdg.configFile."k9s/skins/everforest-light.yaml".text = ''
    # -----------------------------------------------------------------------------
    # Everforest Light
    # https://github.com/sainnhe/everforest/blob/master/palette.md#dark
    # -----------------------------------------------------------------------------
    #
    text: &text "#5c6a72"
    base: &base "#f2efdf"
    overlay: &overlay "#fffbef"
    muted: &muted "#edeada"
    red: &red "#f85552"
    blue: &blue "#3a94c5"
    yellow: &yellow "#dfa000"
    green: &green "#35a77c"
    pink: &pink "#df69ba"
    orange: &orange "#f57d26"

    # Skin...
    k9s:
      # General K9s styles
      body:
        fgColor: *text
        bgColor: *base
        logoColor: *green
      # Command prompt styles
      prompt:
        fgColor: *text
        bgColor: *base
        suggestColor: *green
        border:
          command: *orange
          default: *blue
      # ClusterInfoView styles.
      info:
        fgColor: *green
        sectionColor: *text
      # Dialog styles.
      dialog:
        fgColor: *text
        bgColor: *base
        buttonFgColor: *text
        buttonBgColor: *green
        buttonFocusFgColor: *yellow
        buttonFocusBgColor: *green
        labelFgColor: *yellow
        fieldFgColor: *text
      frame:
        # Borders styles.
        border:
          fgColor: *overlay
          focusColor: *overlay
        menu:
          fgColor: *text
          keyColor: *green
          # Used for favorite namespaces
          numKeyColor: *green
        # CrumbView attributes for history navigation.
        crumbs:
          fgColor: *text
          bgColor: *overlay
          activeColor: *overlay
        # Resource status and update styles
        status:
          newColor: *green
          modifyColor: *red
          addColor: *blue
          errorColor: *pink
          highlightcolor: *yellow
          killColor: *muted
          completedColor: *muted
        # Border title styles.
        title:
          fgColor: *text
          bgColor: *overlay
          highlightColor: *yellow
          counterColor: *green
          filterColor: *green
      views:
        # Charts skins...
        charts:
          bgColor: default
          defaultDialColors:
            - *green
            - *pink
          defaultChartColors:
            - *green
            - *pink
        # TableView attributes.
        table:
          fgColor: *text
          bgColor: *base
          # Header row styles.
          header:
            fgColor: *text
            bgColor: *base
            sorterColor: *red
        # Xray view attributes.
        xray:
          fgColor: *text
          bgColor: *base
          cursorColor: *overlay
          graphicColor: *green
          showIcons: false
        # YAML info styles.
        yaml:
          keyColor: *green
          colonColor: *green
          valueColor: *text
        # Logs styles.
        logs:
          fgColor: *text
          bgColor: *base
          indicator:
            fgColor: *text
            bgColor: *base
  '';

  xdg.configFile."k9s/skins/nord-light.yaml".text = ''
    # -----------------------------------------------------------------------------
    # Nord Light Theme for k9s
    # Based on Nord theme (https://www.nordtheme.com/)
    # -----------------------------------------------------------------------------

    # Nord Light Palette
    text: &text "#2e3440"           # nord0 - main text
    base: &base "#eceff4"           # nord6 - background
    overlay: &overlay "#e5e9f0"     # nord5 - lighter overlay
    muted: &muted "#d8dee9"         # nord4 - selection/muted
    subtle: &subtle "#4c566a"       # nord3 - subtle text

    # Nord Aurora & Frost colors
    red: &red "#bf616a"             # nord11 - aurora red
    orange: &orange "#d08770"       # nord12 - aurora orange
    yellow: &yellow "#ebcb8b"       # nord13 - aurora yellow
    green: &green "#a3be8c"         # nord14 - aurora green
    purple: &purple "#b48ead"       # nord15 - aurora purple
    cyan: &cyan "#88c0d0"           # nord7 - frost cyan
    blue: &blue "#5e81ac"           # nord10 - frost blue
    lightblue: &lightblue "#81a1c1" # nord9 - frost light blue

    # k9s skin configuration
    k9s:
      # General K9s styles
      body:
        fgColor: *text
        bgColor: *base
        logoColor: *blue

      # Command prompt styles
      prompt:
        fgColor: *text
        bgColor: *base
        suggestColor: *green
        border:
          command: *yellow
          default: *lightblue

      # ClusterInfoView styles
      info:
        fgColor: *blue
        sectionColor: *text

      # Dialog styles
      dialog:
        fgColor: *text
        bgColor: *base
        buttonFgColor: *base
        buttonBgColor: *blue
        buttonFocusFgColor: *base
        buttonFocusBgColor: *lightblue
        labelFgColor: *blue
        fieldFgColor: *text

      frame:
        # Borders styles
        border:
          fgColor: *muted
          focusColor: *lightblue

        menu:
          fgColor: *text
          keyColor: *blue
          # Used for favorite namespaces
          numKeyColor: *lightblue

        # CrumbView attributes for history navigation
        crumbs:
          fgColor: *text
          bgColor: *overlay
          activeColor: *lightblue

        # Resource status and update styles
        status:
          newColor: *green
          modifyColor: *yellow
          addColor: *cyan
          errorColor: *red
          highlightcolor: *purple
          killColor: *subtle
          completedColor: *subtle

        # Border title styles
        title:
          fgColor: *text
          bgColor: *overlay
          highlightColor: *blue
          counterColor: *green
          filterColor: *cyan

      views:
        # Charts skins
        charts:
          bgColor: default
          defaultDialColors:
            - *blue
            - *lightblue
            - *cyan
            - *green
          defaultChartColors:
            - *blue
            - *lightblue
            - *cyan
            - *green

        # TableView attributes
        table:
          fgColor: *text
          bgColor: *base
          # Header row styles
          header:
            fgColor: *text
            bgColor: *overlay
            sorterColor: *blue

        # Xray view attributes
        xray:
          fgColor: *text
          bgColor: *base
          cursorColor: *muted
          graphicColor: *blue
          showIcons: false

        # YAML info styles
        yaml:
          keyColor: *blue
          colonColor: *lightblue
          valueColor: *text

        # Logs styles
        logs:
          fgColor: *text
          bgColor: *base
          indicator:
            fgColor: *text
            bgColor: *muted
  '';
}
