Feature: register adventurer
    In order to join in adventures
    As a guest
    I need to register as an adventurer

    Scenario: Succesfully sign up
        Given there is no adventurer with the name of Dash Bo Der
        And I am at the start screen
        When I enter the command of register
        And I enter the name of Dash Bo Der
        And I enter the password of knock.knock...knock
        And I enter the command of save
        Then I should see a note stating "Welcome, fellow mnstr Dash Bo Der!"

    Scenario: Failed with duplicate user
        Given there is an adventurer with the name of Dash Bo Der
        And I am at the start screen
        When I enter the command of register
        And I enter the name of Dash Bo Der
        And I enter the password of knock.knock...knock
        And I enter the command of save
        Then I should see a note stating "It seems like there is an imposter about!"
