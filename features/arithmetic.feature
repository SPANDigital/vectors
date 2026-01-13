Feature: Vector arithmetic
  As a developer
  I want to add and subtract vectors
  So that I can perform basic vector operations

  Scenario: Add two positive vectors
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 4.0 |
      | 5.0 |
      | 6.0 |
    When I add the vectors
    Then the result should be vector
      | 5.0 |
      | 7.0 |
      | 9.0 |

  Scenario: Subtract two positive vectors
    Given vector a is
      | 5.0 |
      | 7.0 |
      | 9.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    When I subtract vector b from vector a
    Then the result should be vector
      | 4.0 |
      | 5.0 |
      | 6.0 |

  Scenario: Add vectors with negative components
    Given vector a is
      | -1.0 |
      | 2.0 |
    And vector b is
      | 3.0 |
      | -4.0 |
    When I add the vectors
    Then the result should be vector
      | 2.0 |
      | -2.0 |

  Scenario: Subtract vectors with negative components
    Given vector a is
      | 3.0 |
      | -4.0 |
    And vector b is
      | -1.0 |
      | 2.0 |
    When I subtract vector b from vector a
    Then the result should be vector
      | 4.0 |
      | -6.0 |

  Scenario: Add zero vectors
    Given vector a is
      | 0.0 |
      | 0.0 |
    And vector b is
      | 0.0 |
      | 0.0 |
    When I add the vectors
    Then the result should be vector
      | 0.0 |
      | 0.0 |

  Scenario: Add vector to zero vector (identity)
    Given vector a is
      | 3.0 |
      | 4.0 |
    And vector b is
      | 0.0 |
      | 0.0 |
    When I add the vectors
    Then the result should be vector
      | 3.0 |
      | 4.0 |

  Scenario: Subtract zero vector (identity)
    Given vector a is
      | 3.0 |
      | 4.0 |
    And vector b is
      | 0.0 |
      | 0.0 |
    When I subtract vector b from vector a
    Then the result should be vector
      | 3.0 |
      | 4.0 |

  Scenario: Subtract vector from itself (returns zero vector)
    Given vector a is
      | 3.0 |
      | 4.0 |
    And vector b is
      | 3.0 |
      | 4.0 |
    When I subtract vector b from vector a
    Then the result should be vector
      | 0.0 |
      | 0.0 |

  Scenario: Add vectors with different lengths
    Given vector a is
      | 1.0 |
      | 2.0 |
    And vector b is
      | 1.0 |
    When I add the vectors
    Then the result should be an error

  Scenario: Subtract vectors with different lengths
    Given vector a is
      | 1.0 |
      | 2.0 |
      | 3.0 |
    And vector b is
      | 1.0 |
      | 2.0 |
    When I subtract vector b from vector a
    Then the result should be an error

  Scenario: Add single component vectors
    Given vector a is
      | 5.0 |
    And vector b is
      | 3.0 |
    When I add the vectors
    Then the result should be vector
      | 8.0 |

  Scenario: Subtract single component vectors
    Given vector a is
      | 5.0 |
    And vector b is
      | 3.0 |
    When I subtract vector b from vector a
    Then the result should be vector
      | 2.0 |
