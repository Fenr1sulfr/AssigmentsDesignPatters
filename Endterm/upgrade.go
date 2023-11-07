package main

type NzothWeaponDecor struct {
	weapon IWeapon
}

func (nd *NzothWeaponDecor) setName(name string) {
	nd.weapon.setName(name + ", N'Zoth")
}
func (nd *NzothWeaponDecor) setDamage(damage int) {
	nd.weapon.setDamage(damage + 10)
}
func (nd *NzothWeaponDecor) getName() string {
	return nd.weapon.getName()
}
func (nd *NzothWeaponDecor) getDamage() int {
	return nd.weapon.getDamage()
}
func (nd *NzothWeaponDecor) getAttack() string {
	return nd.weapon.getAttack() + " , N'Zoth helps you, but hurts your mind"
}

type AzerothDecor struct {
	weapon IWeapon
}

func (az *AzerothDecor) setName(name string) {
	az.weapon.setName(name + ", Consecrated by Azeroth")
}
func (az *AzerothDecor) setDamage(damage int) {
	az.weapon.setDamage(damage + 8)
}
func (az *AzerothDecor) getName() string {
	return az.weapon.getName()
}
func (az *AzerothDecor) getDamage() int {
	return az.weapon.getDamage()
}
func (az *AzerothDecor) getAttack() string {
	return az.weapon.getAttack() + " , Azeroth give you more power"
}
