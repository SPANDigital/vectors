Feature: Check if vector is normalized
  As a developer
  I want to check if vectors are normalized
  So that I can ensure they have a unit length

  Scenario: Check a normalized vector
    Given vector a is
      | 0.6 |
      | 0.8 |
    When I check if vector a is normalized
    Then the result should be true

  Scenario: Check a non-normalized vector
    Given vector a is
      | 3.0 |
      | 4.0 |
    When I check if vector a is normalized
    Then the result should be false

  Scenario: Check a zero vector
    Given vector a is
      | 0.0 |
      | 0.0 |
    When I check if vector a is normalized
    Then the result should be false

  Scenario: Check a single component normalized vector
    Given vector a is
      | 1.0 |
    When I check if vector a is normalized
    Then the result should be true

  Scenario: Check a single component non-normalized vector
    Given vector a is
      | 5.0 |
    When I check if vector a is normalized
    Then the result should be false
