package main

import "fmt"

type LinkedIn interface {
	register()
	notifyUsers()
}
type job struct {
	observerList[] observer
	name string
	skill string
	available bool
}
func newJob(name string,skill string) *job{
	return &job{
		name: name,
		skill: skill,
	}
}
func (j *job) updateJobs(){
	fmt.Printf("%s on %s available\n",j.name,j.skill)
	j.available = true
	j.notifyUsers()
}
func (j *job) register(o observer){
	j.observerList = append(j.observerList, o)
}
func (j *job) notifyUsers() {
	for _, observer := range j.observerList {
		observer.update(j.name,j.skill)
	}
}
type observer interface{
	update(name string,skill string)
	getID() string
}
type customer struct {
	name string
	skill string
}
func (c *customer) update(jobName string,skill string)  {
	if c.skill == skill {
		fmt.Printf("Sending email to  %s for job %s on  \n", c.name, jobName)
	}
}

func (c *customer) getID() string {
	return c.name
}

func main() {

	javaJob := newJob("BackEnd","Java")
	cSharpJob := newJob("BackEnd","C#")
	observerFirst := &customer{name: "Amiran",skill: "Java"}
	observerSecond := &customer{name: "Miras",skill:"C#"}

	//observerThird := &customer{name: "Anton",skill:"Russian"}

	javaJob.register(observerFirst)
	javaJob.register(observerSecond)
	cSharpJob.register(observerFirst)
	cSharpJob.register(observerSecond)

	javaJob.updateJobs()
	cSharpJob.updateJobs()


}
