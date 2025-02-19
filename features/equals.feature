Feature: Check if vectors are equal
  As a developer
  I want to check if two vectors are equal
  So that I can ensure they have the same components

  Scenario: Check equal vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
    When I check if vector a equals vector b
    Then the result should be true

  Scenario: Check non-equal vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 2.0 |
      | 3.0 |
    When I check if vector a equals vector b
    Then the result should be false

  Scenario: Check vectors with different lengths
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    When I check if vector a equals vector b
    Then the result should be false
