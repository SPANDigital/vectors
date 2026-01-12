Feature: Calculate taxicab distance
  As a developer
  I want to calculate the taxicab distance between vectors
  So that I can measure the Manhattan distance between points

  Scenario: Calculate taxicab distance between two 2D vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 4.0 |
      | 6.0 |
    When I calculate the taxicab distance
    Then the result should be 7.0

  Scenario: Calculate taxicab distance between identical vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I calculate the taxicab distance
    Then the result should be 0.0

  Scenario: Calculate taxicab distance between 3D vectors
    Given vector a is
      | 0.0 |
      | 0.0 |
      | 0.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I calculate the taxicab distance
    Then the result should be 6.0

  Scenario: Calculate taxicab distance with negative components
    Given vector a is
      | -1.0 |
      | -2.0 |
    And vector b is
      | 2.0 |
      | 3.0 |
    When I calculate the taxicab distance
    Then the result should be 8.0

  Scenario: Calculate taxicab distance with different length vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    When I calculate the taxicab distance
    Then the result should be an error
