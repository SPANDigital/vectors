Feature: Calculate Euclidean distance
  As a developer
  I want to calculate the Euclidean distance between vectors
  So that I can measure the straight-line distance between points

  Scenario: Calculate Euclidean distance between two 2D vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 4.0 |
      | 6.0 |
    When I calculate the Euclidean distance
    Then the result should be 5.0

  Scenario: Calculate Euclidean distance between identical vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I calculate the Euclidean distance
    Then the result should be 0.0

  Scenario: Calculate Euclidean distance between 3D vectors
    Given vector a is
      | 0.0 |
      | 0.0 |
      | 0.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 2.0 |
    When I calculate the Euclidean distance
    Then the result should be 3.0

  Scenario: Calculate Euclidean distance with negative components
    Given vector a is
      | -1.0 |
      | -1.0 |
    And vector b is
      | 2.0 |
      | 3.0 |
    When I calculate the Euclidean distance
    Then the result should be 5.0

  Scenario: Calculate Euclidean distance with different length vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    When I calculate the Euclidean distance
    Then the result should be an error
