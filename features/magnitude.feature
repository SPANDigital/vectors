Feature: Calculate vector magnitude
  As a developer
  I want to calculate the magnitude of vectors
  So that I can determine their length

  Scenario: Calculate magnitude of a 2D vector
    Given vector a is
      | 3.0 |
      | 4.0 |
    When I calculate the magnitude of vector a
    Then the result should be 5.0

  Scenario: Calculate magnitude of a zero vector
    Given vector a is
      | 0.0 |
      | 0.0 |
    When I calculate the magnitude of vector a
    Then the result should be 0.0

  Scenario: Calculate magnitude of a single component vector
    Given vector a is
      | 5.0 |
    When I calculate the magnitude of vector a
    Then the result should be 5.0

  Scenario: Calculate magnitude of a 3D vector
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 2.0 |
    When I calculate the magnitude of vector a
    Then the result should be 3.0

  Scenario: Calculate magnitude of a negative component vector
    Given vector a is
      | -3.0 |
      | -4.0 |
    When I calculate the magnitude of vector a
    Then the result should be 5.0
