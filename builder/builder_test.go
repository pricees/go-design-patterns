package builder

import "testing"

func TestRobotEngineer(t *testing.T) {
	var oldStyleRobot OldRobotBuilder = BuildOldRobotBuilder()
	robotEngineer := BuildRobotEngineer(oldStyleRobot)
	robotEngineer.MakeRobot()
	var firstRobot *Robot = robotEngineer.GetRobot()

	exp := "Tin Head"
	if firstRobot.robotHead != exp {
		t.Error("expected: ", exp, ", got: ", firstRobot.robotHead)
	}

	exp = "Tin Torso"
	if firstRobot.robotTorso != exp {
		t.Error("expected: ", exp, ", got: ", firstRobot.robotTorso)
	}

	exp = "Rollar Skates"
	if firstRobot.robotLegs != exp {
		t.Error("expected: ", exp, ", got: ", firstRobot.robotLegs)
	}

	exp = "Blowtorch Arms"
	if firstRobot.robotArms != exp {
		t.Error("expected: ", exp, ", got: ", firstRobot.robotArms)
	}
}
