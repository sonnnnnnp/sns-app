import { Injectable } from '@nestjs/common';
import { PrismaService } from './common/prisma/prisma.service';

@Injectable()
export class AppService {
	constructor(private readonly prisma: PrismaService) {}

	async getMessage(): Promise<string> {
		const messages = await this.prisma.message.findUnique({ where: { id: 1 } });
		return messages.body;
	}
}
