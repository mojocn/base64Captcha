# Captcha Clarity Improvements for Small Dimensions

This document describes the improvements made to enhance captcha clarity for small dimensions, specifically addressing issues with 120x30 pixel captchas.

## Issues Fixed

### 1. Font Size Too Small (验证码显示不够清晰)
**Problem**: For small captchas like 120x30, the original font size calculation `height * (rand.Intn(7) + 7) / 16` could result in fonts as small as 13 pixels, making text hard to read.

**Solution**: Added configurable `MinFontSize` and `MaxFontSize` parameters with intelligent defaults:
- For height ≤ 40px: minimum font size is automatically set to 16px
- Default: MinFontSize = 60% of height, MaxFontSize = 80% of height
- Configurable per-driver for custom requirements

### 2. Text Positioning Near Borders (文字开始写的地方太容易在边界)
**Problem**: The original positioning calculation `x := fontWidth*i + fontWidth/fontSize` could place text too close to image borders.

**Solution**: Implemented proper margin-based positioning:
- Added 10% margins on left and right sides
- Centered text within available space
- Ensured text never extends beyond margin boundaries

### 3. Lack of Bold Text Support (希望增加文字加粗的效果)
**Problem**: No option for bold text rendering to improve clarity.

**Solution**: Added `Bold` option to `DriverString`:
- When enabled, renders text with additional stroke offsets
- Configurable per-driver instance
- Backward compatible (default: false)

## New API Features

### DriverString Enhancements

```go
type DriverString struct {
    // ... existing fields ...
    
    // New fields for improved clarity
    MinFontSize int  // Minimum font size (auto-calculated if 0)
    MaxFontSize int  // Maximum font size (auto-calculated if 0)  
    Bold        bool // Enable bold text rendering
}
```

### Usage Examples

#### Basic Usage (Backward Compatible)
```go
// Existing code continues to work unchanged
driver := base64Captcha.NewDriverString(30, 120, 0, 0, 4, "1234567890abcdefghjklmnpqrstuvwxyz", 
    &color.RGBA{255, 255, 255, 0}, nil, []string{})
```

#### Enhanced Small Captcha
```go
// Create driver for small captcha with improved clarity
driver := base64Captcha.NewDriverString(30, 120, 0, 0, 4, "1234567890abcdefghjklmnpqrstuvwxyz", 
    &color.RGBA{255, 255, 255, 0}, nil, []string{})

// Enable improvements
driver.Bold = true         // Better visibility
driver.MinFontSize = 18    // Ensure readability
driver.MaxFontSize = 24    // Control size variation

driver = driver.ConvertFonts()
captcha := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
```

#### Auto-Configuration for Small Captchas
```go
// For height ≤ 40px, font sizes are automatically optimized
driver := base64Captcha.NewDriverString(25, 100, 0, 0, 4, "1234567890", nil, nil, []string{})
driver.Bold = true  // Just enable bold, sizes auto-configured
driver = driver.ConvertFonts()

// MinFontSize will be automatically set to 16 (minimum readable)
// MaxFontSize will be automatically set to 20 (80% of height)
```

## Performance Impact

The improvements have minimal performance impact:
- Font size calculation: O(1) - simple arithmetic
- Bold rendering: ~6x additional character draws (when enabled)
- Positioning: O(1) - improved margin calculations

## Backward Compatibility

✅ **Fully backward compatible**
- All existing code continues to work unchanged
- New fields have sensible defaults
- Original `drawText` method unchanged (calls new method internally)

## Testing

Added comprehensive tests:
- `TestItemChar_drawTextWithFontSize`: Tests new font sizing functionality
- `TestDriverString_SmallCaptcha_ImprovementsFor120x30`: Integration test for 120x30 improvements
- All existing tests continue to pass

## Visual Comparison

The improvements provide significant clarity enhancement for small captchas:

- **Before**: Font sizes 13-24px, text near borders, no bold option
- **After**: Configurable font sizes (default 18-24px), proper margins, optional bold text

Generated test images show clear improvement in readability for 120x30 pixel captchas.