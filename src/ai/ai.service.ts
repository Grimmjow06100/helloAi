import { Inject, Injectable } from '@nestjs/common';
import OpenAI from 'openai';

@Injectable()
export class AiService {
  constructor(@Inject('OPEN_AI_SERVICE') private client: OpenAI) {}

  public async helloAi() {
    const response = await this.client.responses.create({
      model: 'gpt-5.4',
      input: 'hello ai , how are you doing.',
    });
    return response.output_text;
  }
}
