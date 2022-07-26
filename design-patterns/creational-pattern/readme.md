# Creational Design Patterns

## Builder

- Encapsulates an object's construction process along with specifying the various parts of a complex API.
- Enables flexible creation of an object that can have many different representations
- Increases code readability for complex type
- This pattern is used when objects that have complex APIs, multiple constructer options, and several different possible representations.

## Factory

- Allows for the construction of objects when the types of those objects is not predetermined at runtime.
- Produces code that is more readable when there are multiple ways of creating a particular object.
- Situations when object creation needs to be flexible and can't be known beforehand.

## Singleton

- Restricts the instantiation of a class to a single instance and provide global access.
- Allows for lazy initialization of the class
- Situations where you want to ensure there is only one instance of a class: logging, configuration, telemetry, debugging
