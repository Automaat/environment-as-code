# Monitor calibration

Calibrite PROFILER with a Calibrite Display Pro HL, for both external monitors. Last done 2026-10-30.

| Monitor | Panel | Calibrite display type | OSD color mode | Profile |
|---|---|---|---|---|
| ASUS ProArt PA329CRV (primary) | IPS, wide gamut (~98% DCI-P3) | **PFS Phosphor** | **User Mode 1** | `PA329CRV_26-10-30-2.icc` |
| Dell U2515H (rotated) | IPS, sRGB | **White LED** | **Custom Color** | `DELL_U2515H_26-10-30.icc` |

## Files

| Path | What |
|---|---|
| `dotfiles/colorsync/*.icc` | profiles, installed into `~/Library/ColorSync/Profiles` |
| `dotfiles/swift/display-profiles` | display UUID → profile table; the `display color profiles assigned` command applies it |
| `dotfiles/calibrite/userPresets.json` | PROFILER saved presets, linked into the app's support dir |

After recalibrating: copy the new `.icc` into `dotfiles/colorsync/`, point the display's UUIDs at it in `display-profiles` (each Mac gives a monitor its own UUID; `swift ~/.local/bin/display-profiles list` prints this Mac's), drop the old file, then `zakwas apply`. Installed profiles are locked, so the new one needs a new name.

## Calibrite settings

Saved preset: **D65_250nit_G2.2_v2_matrix_211**. **D65_160nit_G2.2_v2_matrix_211** is the same at 160 cd/m² for a dim room.

| Setting | Value |
|---|---|
| Mode | Monitor → Advanced |
| White point | CIE D65 |
| Luminance | 250 cd/m² (same value on both monitors) |
| Contrast | Native |
| Gamma | 2.2 |
| Chromatic adaptation | Bradford |
| ICC profile | Version 2, matrix based |
| Patch set | Advanced, 211 patches |

The display type is not part of the preset. Set it in the dropdown under each monitor on the "Display and Technology Type" screen.

## ASUS PA329CRV OSD

| Menu | Setting | Value |
|---|---|---|
| ProArt Preset | | User Mode 1 (full native gamut; the sRGB and DCI-P3 presets clamp it) |
| ProArt Palette | Brightness | set during calibration to reach the luminance target |
| ProArt Palette | Contrast, Saturation, Hue, Black Level, six-axis | defaults |
| ProArt Palette | Color Temp. | 6500K |
| ProArt Palette | Gamma | 2.2 |
| Image | HDR, Dynamic Dimming, Uniformity Compensation | Off |
| Image | Input Range | Full |
| System Setup | Ambient light and proximity sensors, Eco mode | Off |

## Dell U2515H OSD

| Menu | Setting | Value |
|---|---|---|
| Color | Reset Color Settings | run first |
| Color → Preset Modes | | Custom Color |
| Color → Custom Color | R/G/B gain | 100, unless Calibrite asks for a white point adjustment |
| Color | Input Color Format | RGB |
| Brightness/Contrast | Contrast | 75 |
| Brightness/Contrast | Brightness | set during calibration; 250 cd/m² needs close to maximum |
| Display / Others | Dynamic Contrast, Energy Smart, Uniformity (if present) | Off |

## macOS

- Displays → each monitor: High Dynamic Range off.
- Night Shift and True Tone off.

## Steps

1. Warm the monitors up for 30 minutes. Set the OSD as above.
2. Calibrite: Monitor → Advanced, select the monitor, set its display type.
3. Pick **SAVED Preset** → the 250 nit preset. Check the progress bar shows D65, 250 cd/m², 2.2, Version 2, Matrix, Advanced 211.
4. Start Measurement. Adjust OSD brightness when asked until the reading hits the target.
5. Repeat for the other monitor with the same preset.
6. Leave every OSD setting alone afterwards. Changing brightness or color mode invalidates the profile.

## Check a profile

A correct ASUS run measures its full green primary at about xy (0.28, 0.66), matching the factory green (0.284, 0.665) in the monitor's EDID profile (`PA329CRV-<uuid>.icc` in `~/Library/ColorSync/Profiles`). A green near (0.30, 0.62) means the OSD was in a clamped mode or the display type is wrong. The measured patches are in the profile's `CxF ` tag: zlib-compressed CxF XML with raw XYZ values.

## Gotchas

- A loaded preset overrides edits: changing luminance with a preset selected snaps back. Edit the values in the Progress & Review bar, then save under a new name.
- Changing the OSD can make macOS reconnect the monitor, which resets PROFILER to its start screen and drops unsaved settings. Set the OSD first.
- Factory presets are often 250–350 cd/m² and oversaturated, so an accurate profile at 120–160 cd/m² looks dim and muted at first.
