Feature: Normalize vector
  As a developer
  I want to normalize vectors
  So that I can ensure they have a unit length

  Scenario: Normalize a non-zero vector
    Given vector a is
      | 3.0 |
      | 4.0 |
    When I normalize vector a
    Then the result should be vector
      | 0.6 |
      | 0.8 |

  Scenario: Normalize a zero vector
    Given vector a is
      | 0.0 |
      | 0.0 |
    When I normalize vector a
    Then the result should be vector
      | 0.0 |
      | 0.0 |

  Scenario: Normalize a single component vector
    Given vector a is
      | 5.0 |
    When I normalize vector a
    Then the result should be vector
      | 1.0 |
