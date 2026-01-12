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

  Scenario: Check vectors equal within epsilon tolerance
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0000001 |
      | 2.0000001 |
    And epsilon is 0.001
    When I check if vector a equals vector b with epsilon
    Then the result should be true

  Scenario: Check vectors not equal outside epsilon tolerance
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.1 |
      | 2.1 |
    And epsilon is 0.01
    When I check if vector a equals vector b with epsilon
    Then the result should be false

  Scenario: Check vectors with epsilon at boundary
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.005 |
      | 2.0 |
    And epsilon is 0.01
    When I check if vector a equals vector b with epsilon
    Then the result should be true

  Scenario: Check vectors with different lengths and epsilon
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    And epsilon is 0.001
    When I check if vector a equals vector b with epsilon
    Then the result should be false

  Scenario: Check equal vectors with zero epsilon
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
    And epsilon is 0.0
    When I check if vector a equals vector b with epsilon
    Then the result should be true

  Scenario: Check vectors using DefaultEpsilon constant
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0000000001 |
      | 2.0000000001 |
    And epsilon is 1e-9
    When I check if vector a equals vector b with epsilon
    Then the result should be true
