import { Resend } from 'resend';

export function sendEmail({from , to , subject, html}) {

const resend = new Resend('re_ZsPRRebd_MMf2ArsA9HXEh2yZCu5vY1ed');


resend.emails.send({
  from:  from || 'onboarding@resend.dev',
  to: to ||'admin@mosque.icu',
  subject: subject ||  'Hello World',
  html: html || '<p>Congrats on sending your <strong>first email</strong>!</p>'
});


}
