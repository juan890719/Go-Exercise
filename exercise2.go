package main

import "fmt"

type perso struct {
	name   string
	age    uint
	groups []string
	crush  *perso
}

func (p *perso) getName() string {
	return p.name
}

func (p *perso) getGroups() []string {
	return p.groups
}

func (p *perso) setCrush(crush *perso) {
	p.crush = crush
}

func (p *perso) getCrushName() string {
	if p.crush == nil {
		return "沒有"
	}
	return p.crush.name
}

func printPerso() {
	sa := &perso{
		name: "sakura",
	}

	to := &perso{
		name: "tomoyo",
	}

	tou := &perso{
		name: "touya",
	}

	yu := &perso{
		name: "yu",
	}

	na := &perso{
		name: "naoko",
	}

	sa.setCrush(tou)
	to.setCrush(sa)
	tou.setCrush(yu)
	yu.setCrush(tou)

	fmt.Println("小櫻的本名是?", sa.getName())
	fmt.Println("小櫻喜歡的人是?", sa.getCrushName())
	fmt.Println("Hello", na.name)
	fmt.Println("Nakok喜歡的人是?", na.getCrushName())
}

// 空介面，泛型generic（不區分型態）
func myPrint(i interface{}) {
	fmt.Println(i)
}
