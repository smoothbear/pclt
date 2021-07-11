Feature: Create spring intializr project
  Scenarios: User has chosen all choices
    Then: Request to spring initializr
  Scenarios: User doesn't have chosen all choices
    Then: Default Value
  Scenarios: User entered invalid values
    Then: Print error traceback

Feature: Create github template project
  Scenarios: User has not entered answers
    Then: Git clone
