Feature: Calculate negative inner product
  As a developer
  I want to calculate the negative inner product between vectors
  So that I can use it as a distance metric

  Scenario: Calculate negative inner product of two positive vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 4.0 |
      | 5.0 |
      | 6.0 |
    When I calculate the negative inner product
    Then the result should be -32.0

  Scenario: Calculate negative inner product of orthogonal vectors
    Given vector a is
      | 1.0 |
      | 0.0 |
    And vector b is
      | 0.0 |
      | 1.0 |
    When I calculate the negative inner product
    Then the result should be 0.0

  Scenario: Calculate negative inner product of identical vectors
    Given vector a is
      | 2.0 |
      | 3.0 |
    And vector b is
      | 2.0 |
      | 3.0 |
    When I calculate the negative inner product
    Then the result should be -13.0

  Scenario: Calculate negative inner product with negative components
    Given vector a is
      | -1.0 |
      | 2.0 |
    And vector b is
      | 3.0 |
      | -4.0 |
    When I calculate the negative inner product
    Then the result should be 11.0

  Scenario: Calculate negative inner product with different length vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    When I calculate the negative inner product
    Then the result should be an error
