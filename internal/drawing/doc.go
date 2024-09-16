// Package drawing provides functionality for drawing rectangles and text on a canvas
//
// The drawing package defines a Container type that acts like a div in html.
// It can nest other containers or text and set their layout through declarations
// familiar of that of css flexbox and padding. You can set a div's background color
// and border, and text's alignment, font size and weight.
//
// The package contains the following files:
//
//   - base.go
//     Defines the base types for the package, such Location and Dimension etc.
//
//   - div.go
//     Definition of the Div struct. A rectangle with a background color and border.
//
//   - container.go
//     Extends the div functionality to allow nesting.
//
//   - text.go
//     Provides the functionality to write text.
//
//   - colour.go
//     Allows new colours to be created.
//
//   - canvas.go
//     Provides the Canvas to draw to and functions to export as an image.Image.
package drawing
