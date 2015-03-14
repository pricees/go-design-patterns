package builder

type RobotPlan interface {
	SetRobotHead(string)
	SetRobotTorso(string)
	SetRobotArms(string)
	SetRobotLegs(string)
}

type Robot struct {
	robotHead  string
	robotTorso string
	robotArms  string
	robotLegs  string
}

func (r *Robot) SetRobotHead(head string) {
	r.robotHead = head
}

func (r *Robot) SetRobotTorso(torso string) {
	r.robotTorso = torso
}

func (r *Robot) SetRobotArms(arms string) {
	r.robotArms = arms
}

func (r *Robot) SetRobotLegs(legs string) {
	r.robotLegs = legs
}

type RobotBuilder interface {
	BuildRobotHead()
	BuildRobotTorso()
	BuildRobotArms()
	BuildRobotLegs()
	GetRobot() *Robot
}

/////////////

type OldRobotBuilder struct {
	robot *Robot
}

func BuildOldRobotBuilder() OldRobotBuilder {
	return OldRobotBuilder{robot: &Robot{}}
}

func (o OldRobotBuilder) BuildRobotHead() {
	o.robot.SetRobotHead("Tin Head")
}

func (o OldRobotBuilder) BuildRobotTorso() {
	o.robot.SetRobotTorso("Tin Torso")
}

func (o OldRobotBuilder) BuildRobotArms() {
	o.robot.SetRobotArms("Blowtorch Arms")
}

func (o OldRobotBuilder) BuildRobotLegs() {
	o.robot.SetRobotLegs("Rollar Skates")
}

func (o OldRobotBuilder) GetRobot() *Robot {
	return o.robot
}

/////////////

type RobotEngineer struct {
	robotBuilder RobotBuilder
}

func BuildRobotEngineer(robotBuilder RobotBuilder) RobotEngineer {
	return RobotEngineer{robotBuilder}
}

func (r RobotEngineer) GetRobot() *Robot {
	return r.robotBuilder.GetRobot()
}

func (r RobotEngineer) MakeRobot() {
	r.robotBuilder.BuildRobotHead()
	r.robotBuilder.BuildRobotTorso()
	r.robotBuilder.BuildRobotArms()
	r.robotBuilder.BuildRobotLegs()
}
