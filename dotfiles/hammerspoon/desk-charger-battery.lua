-- ========================================
-- Desk Charger Battery Controller
-- ========================================
-- Monitors iPhone and AirPods case battery levels
-- via AirBuddy Shortcuts and controls Home Assistant
-- smart plug via webhook
--
-- Installation:
-- 1. Copy to ~/.hammerspoon/desk-charger-battery.lua
-- 2. Add to ~/.hammerspoon/init.lua:
--    require("desk-charger-battery")
-- 3. Create Shortcuts:
--    - "Get iPhone Battery" - returns battery % as number
--    - "Get AirPods Battery" - returns battery % as number
-- 4. Reload Hammerspoon config

local log = hs.logger.new('desk-charger', 'info')

-- ========================================
-- Configuration
-- ========================================
local config = {
    -- Home Assistant webhook URL
    ha_webhook_url = "http://homeassistant.local:8123/api/webhook/desk_charger_battery_a8f3c901_4d2a_4b18_9f7e_8d3c4e7a6b12",

    -- Check interval (seconds)
    check_interval = 300, -- 5 minutes

    -- Shortcut names
    iphone_shortcut = "Get iPhone Battery",
    airpods_shortcut = "Get AirPods Battery",

    -- Battery thresholds
    thresholds = {
        iphone_low = 50,      -- Start charging if iPhone < 50%
        airpods_low = 20,     -- Start charging if AirPods case < 20%
        both_high = 80,       -- Stop charging when both >= 80%
    },
}

-- ========================================
-- State Management
-- ========================================
local state = {
    timer = nil,
    last_iphone_battery = nil,
    last_airpods_battery = nil,
}

-- ========================================
-- Battery Level Retrieval
-- ========================================
local function getBatteryLevel(shortcut_name)
    local success, result = pcall(function()
        return hs.shortcuts.run(shortcut_name)
    end)

    if not success then
        log.e(string.format("Failed to run shortcut '%s': %s", shortcut_name, result))
        return nil
    end

    local battery = tonumber(result)
    if not battery then
        log.e(string.format("Invalid battery value from '%s': %s", shortcut_name, result))
        return nil
    end

    return battery
end

local function getBatteryLevels()
    local iphone = getBatteryLevel(config.iphone_shortcut)
    local airpods = getBatteryLevel(config.airpods_shortcut)

    if iphone then
        state.last_iphone_battery = iphone
    end
    if airpods then
        state.last_airpods_battery = airpods
    end

    return iphone, airpods
end

-- ========================================
-- Home Assistant Webhook
-- ========================================
local function sendWebhook(iphone_battery, airpods_battery, on_success)
    local json = hs.json.encode({
        iphone_battery = iphone_battery,
        airpods_battery = airpods_battery,
    })

    local headers = {
        ["Content-Type"] = "application/json"
    }

    log.d(string.format("Sending webhook: iPhone=%d%%, AirPods=%d%%", iphone_battery, airpods_battery))

    hs.http.asyncPost(
        config.ha_webhook_url,
        json,
        headers,
        function(status, body, headers)
            if status == 200 then
                log.d(string.format("Webhook sent successfully"))
                if on_success then
                    on_success()
                end
            else
                log.e(string.format("Webhook failed (status %d): %s", status or "nil", body or "no response"))
            end
        end
    )
end

-- ========================================
-- Battery Check Handler
-- ========================================
local function checkBatteries()
    local iphone, airpods = getBatteryLevels()

    if not iphone or not airpods then
        log.w("Could not retrieve battery levels, skipping this check")
        return
    end

    log.i(string.format("Battery levels: iPhone=%d%%, AirPods case=%d%%", iphone, airpods))

    -- Send to Home Assistant for automation decision
    sendWebhook(iphone, airpods)
end

-- ========================================
-- Initialization
-- ========================================
local function start()
    if state.timer then
        state.timer:stop()
    end

    log.i("Starting desk charger battery monitor")
    log.i(string.format("Webhook URL: %s", config.ha_webhook_url))
    log.i(string.format("Check interval: %d seconds", config.check_interval))
    log.i(string.format("Thresholds: iPhone<%d%% or AirPods<%d%% → charge ON, both≥%d%% → charge OFF",
        config.thresholds.iphone_low,
        config.thresholds.airpods_low,
        config.thresholds.both_high))

    -- Initial check
    checkBatteries()

    -- Start periodic checks
    state.timer = hs.timer.new(config.check_interval, checkBatteries)
    state.timer:start()

    log.i("Desk charger battery monitor started")
end

local function stop()
    if state.timer then
        state.timer:stop()
        state.timer = nil
    end
    log.i("Desk charger battery monitor stopped")
end

-- ========================================
-- Module Interface
-- ========================================
local module = {
    start = start,
    stop = stop,
    config = config,
}

-- Auto-start on load
start()

return module
