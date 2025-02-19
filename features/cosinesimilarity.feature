Feature: Cosine Similarity

  Scenario: Calculate cosine similarity of two non-zero vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 4.0 |
      | 5.0 |
      | 6.0 |
    When I calculate the cosine similarity
    Then the result should be 0.9746318461970762

  Scenario: Calculate cosine similarity of a vector with itself
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I calculate the cosine similarity
    Then the result should be 1.0

  Scenario: Calculate cosine similarity of orthogonal vectors
    Given vector a is
      | 1.0 |
      | 0.0 |
      | 0.0 |
    And vector b is
      | 0.0 |
      | 1.0 |
      | 0.0 |
    When I calculate the cosine similarity
    Then the result should be 0.0

  Scenario: Calculate cosine similarity with a zero vector
    Given vector a is
      | 0.0 |
      | 0.0 |
      | 0.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I calculate the cosine similarity
    Then the result should be an error
