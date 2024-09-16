// Package page defines the high level components to be drawn.
//
// # The page package makes use of the drawing package to create
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
//     Provides the Canvas to draw on, and functions to export as an image.Image.
package page
