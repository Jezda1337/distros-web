import { Distro, PrismaClient } from "@prisma/client"

const prisma = new PrismaClient()

// create
export async function createDistro(distro: Distro) {
	if (distro.name === "" || distro.name.length === 0)
		throw new Error("Name is required.")

	const newDistro = await prisma.distro.create({
		data: distro,
	})

	return newDistro
}

export async function addToWaitingList(distro: any) {
	console.log("------------------")
	console.log(distro)
	console.log("------------------")

	const newDistro = await prisma.distro.create({
		data: {
			...distro,
			de: {
				create: [...distro.de],
			},
		},
	})

	return newDistro
}

// read
export function getDistros() {
	const distros = prisma.distro.findMany({
		include: {
			de: true,
			screenShoots: true,
		},
	})
	return distros
}

export function getDistroByName(distroName: string) {
	if (distroName.length === 0 || !distroName)
		throw new Error("Name is required")

	const distro = prisma.distro.findFirst({
		where: {
			name: distroName,
		},
		include: {
			de: true,
			screenShoots: true,
		},
	})

	return distro
}

export async function getWaitingDistros() {
	const distros = await prisma.distro.findMany()
	return distros
}

export async function getWaitingDistro(distroName: string) {
	const distro = await prisma.distro.findFirst({
		where: {
			name: distroName,
		},
	})
	return distro
}

// update

// delete
