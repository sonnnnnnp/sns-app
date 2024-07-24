import { PrismaClient, Prisma } from '@prisma/client';

const prisma = new PrismaClient();

const messageData: Prisma.MessageCreateInput[] = [
	{
		body: 'Message from postgres!',
	},
];

const transfer = async () => {
	const messages = [];
	for (const data of messageData) {
		const message = prisma.message.create({
			data: data,
		});
		messages.push(message);
	}
	return await prisma.$transaction(messages);
};

transfer()
	.catch((e) => {
		console.error(e);
		process.exit(1);
	})
	.finally(async () => {
		await prisma.$disconnect();
	});
